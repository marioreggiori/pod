package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "julia"}

	rootCmd.AddCommand(cmd("julia", "julia runtime", opts))
}
