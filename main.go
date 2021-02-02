package main

import (
	"github.com/marioreggiori/pod/cmd"
	"github.com/marioreggiori/pod/store"
)

func main() {
	defer store.Close()
	cmd.Execute()
}
