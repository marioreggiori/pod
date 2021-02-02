package utils

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/marioreggiori/pod/global"
	"github.com/moby/term"
	"golang.org/x/crypto/ssh/terminal"
)

type RunWithDockerOptions struct {
	Image      string
	User       string
	WorkingDir string
	Tag        string
}

func (opts *RunWithDockerOptions) Validate() error {
	missingOption := "Option [%s] is missing!"
	if opts.Image == "" {
		return fmt.Errorf(missingOption, "Image")
	}
	if opts.User == "" {
		opts.User = "1000"
	}
	if opts.WorkingDir == "" {
		opts.WorkingDir = "/usr/src/app"
	}
	return nil
}

func (opts *RunWithDockerOptions) ImageWithTag() string {
	res := opts.Image
	if tagFromFlag := global.ImageTag(); tagFromFlag != "" {
		res += ":" + tagFromFlag
	} else if opts.Tag != "" {
		res += ":" + opts.Tag
	}
	return res
}

func RunWithDocker(cmd []string, opts *RunWithDockerOptions) {
	if err := opts.Validate(); err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, opts.ImageWithTag(), types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	if global.IsVerbose() {
		termFd, isTerm := term.GetFdInfo(os.Stderr)
		jsonmessage.DisplayJSONMessagesStream(reader, os.Stderr, termFd, isTerm, nil)
	} else {
		_, err = ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
	}

	// todo map ports & volumes

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        opts.ImageWithTag(),
		Cmd:          cmd,
		Tty:          true,
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		OpenStdin:    true,
		WorkingDir:   opts.WorkingDir,
		User:         opts.User,
	}, &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: dir,
				Target: opts.WorkingDir,
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	defer cli.ContainerRemove(context.Background(), resp.ID, types.ContainerRemoveOptions{
		Force: true,
	})

	waiter, err := cli.ContainerAttach(ctx, resp.ID, types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})

	if err != nil {
		panic(err)
	}

	go io.Copy(os.Stdout, waiter.Reader)
	go io.Copy(os.Stderr, waiter.Reader)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fd := int(os.Stdin.Fd())
	var oldState *terminal.State
	if terminal.IsTerminal(fd) {
		oldState, err = terminal.MakeRaw(fd)
		if err != nil {
			panic(err)
		}

		go func() {
			for {
				consoleReader := bufio.NewReaderSize(os.Stdin, 1)
				input, _ := consoleReader.ReadByte()
				if false /*input == 3*/ { // ctl-c
					cli.ContainerRemove(context.Background(), resp.ID, types.ContainerRemoveOptions{
						Force: true,
					})
				}
				waiter.Conn.Write([]byte{input})
			}
		}()
	}

	statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	if terminal.IsTerminal(fd) {
		terminal.Restore(fd, oldState)
	}

}
