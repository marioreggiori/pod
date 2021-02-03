package main

import (
	"log"

	"github.com/marioreggiori/pod/cmd"
	"github.com/marioreggiori/pod/global"
	"github.com/marioreggiori/pod/store"
)

func cleanupTheMess() {
	store.Close()
	if global.IsVerbose() {
		return
	}
	if r := recover(); r != nil {
		log.Fatal(r)
	}
}

func main() {
	defer cleanupTheMess()
	cmd.Execute()
}
