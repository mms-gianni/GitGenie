package main

import (
	"github.com/mms-gianni/GitGenie/cmd"
	"github.com/mms-gianni/GitGenie/pkg/genie"
)

func main() {
	config, _ := cmd.Execute()

	genie.Init(config)
	genie.Run()
}
