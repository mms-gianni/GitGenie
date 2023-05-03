package genie

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

var config *Config

type Config struct {
	OpenAiApiHost  string
	OpenAiApiToken string
	Suggestions    string
	Length         string
	Max_tokens     string
	Skipedit       bool
	Language       string
	Signoff        bool
}

func Init(c *Config) {
	config = c
	loadFromLanguageYaml()
}

func Run() {

	initClient()
	diff := diff()

	commitMessages := submitToApiChat(diff)
	commitMessage := selectCommitMessage(commitMessages)
	commitMessage = editCommitMessage(commitMessage)

	commit(commitMessage)
}

func diff() string {
	out, err := exec.Command("git", "diff", "--cached", "-u").Output()
	if err != nil {
		fmt.Println(err)
	}

	if len(out) == 0 {
		fmt.Print("No changes to commit\n\n")

		status()
		os.Exit(1)
	}

	return string(out)
}

func status() {
	//out, err := exec.Command("git", "status", "-s", "-uno").Output()
	out, err := exec.Command("git", "status").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func selectCommitMessage(options []string) string {
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

func editCommitMessage(commitMsg string) string {

	if config.Skipedit && commitMsg != "" {
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

func commit(commitMsg string) {
	signoff := "--no-signoff"
	if config.Signoff {
		signoff = "--signoff"
	}

	out, err := exec.Command("git", "commit", signoff, "-m", commitMsg).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
