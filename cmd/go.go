package cmd

import (
	"github.com/marioreggiori/pod/utils"
)

func init() {
	var opts = &utils.RunWithDockerOptions{
		Image:      "golang",
		User:       "0",
		WorkingDir: "/go/src/app",
	}

	rootCmd.AddCommand(cmd("go", opts))
}
