package cmd

import (
	"context"
	"log"

	"github.com/marioreggiori/pod/global"
	"github.com/marioreggiori/pod/utils"
	"github.com/spf13/cobra"
)

var flags = &global.Flags{}

var rootCmd = &cobra.Command{
	Use:              "pod",
	Short:            "Run your favorite commands using containers",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		flags.Set()
	},
}

func init() {
	rootCmd.Flags().BoolVar(&flags.Verbose, "verbose", false, "display additional output")
	rootCmd.Flags().StringVarP(&flags.ImageTag, "tag", "t", "", "set image tag")
	rootCmd.Flags().StringArrayVarP(&flags.EnvVariables, "env", "e", nil, "set environment variable")
	rootCmd.Flags().StringArrayVarP(&flags.MappedPorts, "port", "p", nil, "map port")
	rootCmd.Flags().StringArrayVarP(&flags.MappedVolumes, "volume", "v", nil, "map volume")
}

func Execute() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}

func cmd(use, desc string, opts *utils.RunWithDockerOptions) *cobra.Command {
	return &cobra.Command{
		Use:                use,
		Short:              desc,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			utils.RunWithDocker(append([]string{use}, args...), opts)
		},
	}
}
