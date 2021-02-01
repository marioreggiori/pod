package cmd

import (
	"context"
	"log"

	"github.com/marioreggiori/pod/global"
	"github.com/marioreggiori/pod/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var verbose bool
var imageTag string
var envVariables []string
var mappedPorts []string
var mappedVolumes []string

var rootCmd = &cobra.Command{
	Use:              "pod",
	Short:            "Run your favorite commands using containers",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		global.SetIsVerbose(verbose)
		global.SetImageTag(imageTag)
	},
}

func init() {
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "display additional output")
	rootCmd.Flags().StringVarP(&imageTag, "tag", "t", "", "set image tag")
	rootCmd.Flags().StringArrayVarP(&envVariables, "env", "e", nil, "set environment variable")
	rootCmd.Flags().StringArrayVarP(&mappedPorts, "port", "p", nil, "map port")
	rootCmd.Flags().StringArrayVarP(&mappedVolumes, "volume", "v", nil, "map volume")

	rootCmd.AddCommand(&cobra.Command{
		Use: "generate docs",
		Run: func(cmd *cobra.Command, args []string) {
			err := doc.GenMarkdownTree(rootCmd, "./docs")
			if err != nil {
				log.Fatal(err)
			}
		},
	})
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
