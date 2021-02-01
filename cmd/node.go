package cmd

import (
	"github.com/marioreggiori/pod/utils"
)

func init() {
	var opts = &utils.RunWithDockerOptions{
		Image:      "node",
		User:       "1000",
		WorkingDir: "/usr/src/app",
	}

	rootCmd.AddCommand(cmd("node", opts))
	rootCmd.AddCommand(cmd("npm", opts))
	rootCmd.AddCommand(cmd("npx", opts))
}
