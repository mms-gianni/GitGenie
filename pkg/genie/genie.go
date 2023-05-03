package genie

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mms-gianni/GitGenie/cmd"
)

var config *Config

type Config struct {
	openAiApiHost  string
	openAiApiToken string
	suggestions    string
	length         string
	max_tokens     string
	skipedit       bool
}

func LoadConfig() {

	config = &Config{}
	config.openAiApiHost = cmd.OpenAiApiHost
	config.openAiApiToken = cmd.OpenAiApiToken
	config.suggestions = cmd.Suggestions
	config.length = cmd.Length
	config.skipedit = cmd.Fast
	config.max_tokens = cmd.MaxTokens

}

func Diff() string {
	out, err := exec.Command("git", "diff", "--cached", "-u").Output()
	if err != nil {
		fmt.Println(err)
	}

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
	msg := ""
	prompt := &survey.Select{
		Message: "Select a commit message:",
		Options: options,
	}
	survey.AskOne(prompt, &msg)

	if msg == "<empty>" {
		msg = ""
	}

	return msg
}

func EditCommitMessage(commitMsg string) string {

	if config.skipedit && commitMsg != "" {
		return commitMsg
	}

	editedCommitMsg := commitMsg
	prompt := &survey.Editor{
		Message:       "Edit commit message:",
		Default:       commitMsg,
		HideDefault:   true,
		AppendDefault: true,
	}
	survey.AskOne(prompt, &editedCommitMsg)

	return editedCommitMsg
}

func Commit(commitMsg string) {
	out, err := exec.Command("git", "commit", "-m", commitMsg).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
