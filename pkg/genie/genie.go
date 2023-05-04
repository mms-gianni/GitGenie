package genie

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v3"
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
	out, err := exec.Command("git", "diff", "--cached", "--no-color", "-U10", "-u", "--ignore-space-at-eol", "--ignore-all-space", "--ignore-submodules").Output()
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

func getGitRoot() string {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimRight(string(out), "\n")
}

func (c *repoConfig) loadRepoConfig(gitRoot string) *repoConfig {
	// read from config file

	yamlFile, err := os.ReadFile(gitRoot + "/.gitgenie.yaml")
	if err != nil {
		c.Loaded = false
		return c
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println("Loaded .gitgenie.yaml file")

	return c
}

type repoConfig struct {
	Loaded      bool   `yaml:"loaded"`
	Lang        string `yaml:"lang"`
	Description string `yaml:"description"`
}
