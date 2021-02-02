package cmd

import (
	"github.com/marioreggiori/pod/utils"
)

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "mongo"}

	rootCmd.AddCommand(cmd("mongo", "MongoDB client", opts))
}
