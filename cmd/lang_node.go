package cmd

import (
	"github.com/marioreggiori/pod/utils"
)

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "node"}

	rootCmd.AddCommand(cmd("node", "Node.js interpreter", opts))
	rootCmd.AddCommand(cmd("npm", "Node.js package manager", opts))
	rootCmd.AddCommand(cmd("npx", "Node.js command-line tool", opts))
}
