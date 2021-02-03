package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "python"}

	rootCmd.AddCommand(cmd("python", "Python interpreter", opts))
	rootCmd.AddCommand(cmd("pip", "Python package manager", opts))
}
