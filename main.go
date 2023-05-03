package main

import (
	"github.com/mms-gianni/GitGenie/cmd"
	"github.com/mms-gianni/GitGenie/pkg/genie"
)

var (
	// Version is the version of the application
	Version string
	// BuildDate is the date the application was built
	BuildDate string
)

func main() {

	config := cmd.Execute()
	genie.Init(config)
	genie.Run()

}
