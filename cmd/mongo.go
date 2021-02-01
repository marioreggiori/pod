package cmd

import (
	"github.com/marioreggiori/pod/utils"
)

func init() {
	var opts = &utils.RunWithDockerOptions{
		Image:      "mongo",
		User:       "0",
		WorkingDir: "/usr/src/app",
	}

	rootCmd.AddCommand(cmd("mongo", "MongoDB client", opts))
}
