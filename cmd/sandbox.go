package cmd

import "github.com/marioreggiori/pod/utils"

func init() {
	var opts = &utils.RunWithDockerOptions{Image: "ubuntu", User: "0", WorkingDir: "/root"}

	rootCmd.AddCommand(cmdWithAlias("sandbox", "bash", "Ubuntu sandbox", opts))
}
