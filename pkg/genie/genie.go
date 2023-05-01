package genie

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

func Diff() string {
	out, err := exec.Command("git", "diff", "-u").Output()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(out))

	if len(out) == 0 {
		fmt.Println("No changes to commit")
		os.Exit(1)
	}

	return string(out)
}

func Status() {
	out, err := exec.Command("git", "status", "-s", "-uno").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func SelectCommitMessage(options []string) string {
	color := ""
	prompt := &survey.Select{
		Message: "Select a commit message:",
		Options: options,
	}
	survey.AskOne(prompt, &color)

	return color
}
