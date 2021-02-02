package utils

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/go-connections/nat"
	"github.com/marioreggiori/pod/global"
	"github.com/moby/term"
	"golang.org/x/crypto/ssh/terminal"
)

type RunWithDockerOptions struct {
	Image               string
	User                string
	WorkingDir          string
	Tag                 string
	DisableWorkdirMount bool
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
	cli, err := client.NewEnvClient()
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

	// todo map ports
	portMap := nat.PortMap{}
	portSet := nat.PortSet{}
	for _, v := range global.Ports() {
		portParts := strings.Split(v, ":")
		var port nat.Port

		binding := []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: portParts[0]}}

		portString := portParts[0]

		if len(portParts) == 2 {
			portString = portParts[1]
		}

		port, err := nat.NewPort("tcp", portString)
		if err != nil {
			panic(err)
		}
		portSet[port] = struct{}{}
		portMap[port] = binding
	}

	mounts := []mount.Mount{}
	if !opts.DisableWorkdirMount {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: dir,
			Target: opts.WorkingDir,
		})
	}

	for _, v := range global.Mounts() {
		mountParts := strings.Split(v, ":")
		if len(mountParts) != 2 {
			panic("invalid mount option")
		}
		localPath, err := filepath.Abs(mountParts[0])
		if err != nil {
			panic(err)
		}

		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: localPath,
			Target: mountParts[1],
		})
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        opts.ImageWithTag(),
		Cmd:          cmd,
		Tty:          true,
		OpenStdin:    true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir:   opts.WorkingDir,
		User:         opts.User,
		ExposedPorts: portSet,
	}, &container.HostConfig{
		AutoRemove:   true,
		Mounts:       mounts,
		PortBindings: portMap,
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
