package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{
		Image:      "openjdk",
		User:       "1000",
		WorkingDir: "/usr/src/app",
	}

	rootCmd.AddCommand(cmd("java", "Java (OpenJDK) interpreter", opts))
	rootCmd.AddCommand(cmd("javac", "Java (OpenJDK) compiler", opts))
}
