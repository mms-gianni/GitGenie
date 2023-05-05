package main

import (
	log "github.com/pieterclaerhout/go-log"

	"github.com/mms-gianni/GitGenie/cmd"
	"github.com/mms-gianni/GitGenie/pkg/genie"
)

func main() {

	config, err := cmd.Execute()
	if err != nil {
		log.Fatalf("failure: %s", err)
	}
	genie.Init(config)
	genie.Run()

}
