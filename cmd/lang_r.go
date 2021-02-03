package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "r-base"}

	rootCmd.AddCommand(cmd("Rscript", "R interpreter", opts))
}
