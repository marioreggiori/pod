package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var customCmd = &cobra.Command{
		Use:   "custom",
		Short: "Manage custom commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			// todo: list custom commands
			fmt.Println("<list of custom commands>")
		},
	}

	var addCustomCmd = &cobra.Command{
		Use:     "add [command] [image] [description]",
		Short:   "Add custom command",
		Example: "pod custom add npm node \"Node.js interpreter\"",
		Args:    cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	var removeCustomCmd = &cobra.Command{
		Use:   "remove [command]",
		Short: "Remove custom command",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	customCmd.AddCommand(addCustomCmd)
	customCmd.AddCommand(removeCustomCmd)

	rootCmd.AddCommand(customCmd)
}
