package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "gcc"}

	rootCmd.AddCommand(cmd("gcc", "GNU Compiler Collection", opts))
	rootCmd.AddCommand(cmd("g++", "GNU Compiler Collection", opts))
}
