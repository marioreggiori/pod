package cmd

import (
	"github.com/marioreggiori/pod/store"
	"github.com/marioreggiori/pod/utils"
	"github.com/spf13/cobra"
)

func init() {
	var customCmd = &cobra.Command{
		Use:   "custom",
		Short: "Manage custom commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	var addCustomCmd = &cobra.Command{
		Use:     "add [command] [image] [description]",
		Short:   "Add custom command",
		Example: "pod custom add npm node \"Node.js package manager\"",
		Args:    cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			err := store.AddCustom(&store.Custom{Command: args[0], Image: args[1], Description: args[2]})
			if err != nil {
				panic(err)
			}
		},
	}

	var removeCustomCmd = &cobra.Command{
		Use:   "remove [command]",
		Short: "Remove custom command",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := store.RemoveCustom(args[0])
			if err != nil {
				panic(err)
			}
		},
	}

	customCmd.AddCommand(addCustomCmd)
	customCmd.AddCommand(removeCustomCmd)

	rootCmd.AddCommand(customCmd)

	for _, v := range store.GetCustom() {
		rootCmd.AddCommand(cmd(v.Command, v.Description, &utils.RunWithDockerOptions{Image: v.Image}))
	}
}
