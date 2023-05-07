package genie

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	log "github.com/pieterclaerhout/go-log"

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
	Debug          bool
	Diffcontext    string
}

func Init(c *Config) {
	config = c
	loadFromLanguageYaml()
	initClient()

	if config.Debug {
		log.DebugMode = true
		fmt.Printf("Config %+v\n", config)
	}
}

func Run() {

	diff := diff()

	commitMessages := submitToApiChat(diff)
	commitMessage := selectCommitMessage(commitMessages)
	commitMessage = editCommitMessage(commitMessage)

	commit(commitMessage)
}

func diff() string {
	out, err := exec.Command("git", "diff", "--cached", "--no-color", "-U"+config.Diffcontext, "-u", "--ignore-space-at-eol", "--ignore-all-space", "--ignore-submodules").Output()
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

	ticket := getTicket()
	if ticket != "" {
		fmt.Println("Found Ticket: " + ticket)
	}

	if config.Skipedit && commitMsg != "" {
		return commitMsg
	}

	editedCommitMsg := commitMsg
	prompt := &survey.Editor{
		Message:       "Edit commit message:",
		Default:       ticket + " " + commitMsg,
		HideDefault:   true,
		AppendDefault: true,
	}
	survey.AskOne(prompt, &editedCommitMsg)

	return editedCommitMsg
}

/*
func getLogs() []string {
	out, err := exec.Command("git", "log", "--pretty=\"format:%s\"", "-n", "5").Output()
	if err != nil {
		fmt.Println(err)
	}

	split := strings.Split(string(out), "\n")
	return split
}
*/

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

func getBranch() string {
	out, err := exec.Command("git", "branch", "--show-current").Output()
	//out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimRight(string(out), "\n")
}

func getTicket() string {
	var branch string = getBranch()
	r, _ := regexp.Compile(`^([A-Z]+-\d+).*`) // Matches only JIRA tickets for now
	matches := r.FindStringSubmatch(branch)
	var ticket string = string(matches[1])
	return ticket
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

	yamlFile, err := os.ReadFile(gitRoot + "/.gitgenie")
	if err != nil {
		c.Loaded = false
		return c
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

type repoConfig struct {
	Loaded      bool   `yaml:"loaded"`
	Language    string `yaml:"language"`
	Description string `yaml:"description"`
}
