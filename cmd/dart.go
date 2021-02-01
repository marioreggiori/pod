package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{
		Image:      "google/dart",
		User:       "1000",
		WorkingDir: "/app",
	}

	rootCmd.AddCommand(cmd("dart", "Dart runtime", opts))
	rootCmd.AddCommand(cmd("pub", "Dart package manager", opts))
}
