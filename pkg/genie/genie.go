package genie

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

var config *Config

type Config struct {
	openAiApiHost  string
	openAiApiToken string
	suggestions    string
	length         string
	max_tokens     string
}

func loadConfig() {
	config = &Config{}
	config.openAiApiHost = getEnv("OPENAI_API_HOST", "api.openai.com")
	config.openAiApiToken = os.Getenv("OPENAI_API_KEY")
	config.suggestions = getEnv("GENIE_SUGESTIONS", "3")
	config.length = getEnv("GENIE_LENGTH", "medium")
	//config.max_tokens = getEnv("GENIE_MAX_TOKENS", "300")

	switch config.length {
	case "short":
		config.max_tokens = "100"
	case "medium":
		config.max_tokens = "300"
	case "long":
		config.max_tokens = "500"
	case "verylong":
		config.max_tokens = "1000"
	default:
		config.max_tokens = "300"
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
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
	color := ""
	prompt := &survey.Select{
		Message: "Select a commit message:",
		Options: options,
	}
	survey.AskOne(prompt, &color)

	if color == "<empty>" {
		color = ""
	}

	return color
}

func EditCommitMessage(commitMsg string) string {
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
