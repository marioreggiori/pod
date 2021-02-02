package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	var docGenCmd = &cobra.Command{
		Use:   "doc-gen",
		Short: "Generate markdown docs",
		Run: func(cmd *cobra.Command, args []string) {
			err := doc.GenMarkdownTree(rootCmd, "./docs")
			if err != nil {
				panic(err)
			}
		},
	}
	rootCmd.AddCommand(docGenCmd)
}
