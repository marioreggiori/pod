package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "rust"}

	rootCmd.AddCommand(cmd("rustc", "Rust compiler", opts))
	rootCmd.AddCommand(cmd("rustdoc", "Rust documentation", opts))
	rootCmd.AddCommand(cmd("cargo", "Rust package manager", opts))
}
