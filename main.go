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

	cmd.Execute()

	//genie.Status()

	genie.LoadConfig()
	genie.InitClient()
	diff := genie.Diff()

	//commitMessages := genie.SubmitToApi(diff)
	commitMessages := genie.SubmitToApiChat(diff)
	commitMessage := genie.SelectCommitMessage(commitMessages)
	commitMessage = genie.EditCommitMessage(commitMessage)

	genie.Commit(commitMessage)

}
