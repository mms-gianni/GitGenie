package genie

import (
	"testing"
)

func TestGetSystem(t *testing.T) {
	Init(&Config{Language: "en"})
	systemPrompt := getSystem("en")
	if systemPrompt != "You are a programmer. Generate a Git commit messages based on the provided changes." {
		t.Errorf("System could not load SystemPrompt %s", systemPrompt)
	}
}

func TestGetUser(t *testing.T) {
	Init(&Config{Language: "en"})
	userPrompt := getUser("en")
	if userPrompt != "Describe the code changes and the purpose in one english sentence. Start with a verb." {
		t.Errorf("User could not load UserPrompt %s", userPrompt)
	}
}

func TestGetGitRoot(t *testing.T) {
	gitRoot := getGitRoot()
	if len(gitRoot) < 2 {
		t.Errorf("gitRoot is to short %s", gitRoot)
	}
}

func TestLoadRepoConfig(t *testing.T) {
	var gitRoot = getGitRoot()
	var repoConfig repoConfig
	repoConfig.loadRepoConfig(gitRoot)
	if repoConfig.Language != "en" {
		t.Errorf("repoConfig.Language is not en %s", repoConfig.Language)
	}
}

func TestEditCommitMessage(t *testing.T) {
	var res = editCommitMessage("test")
	if res != "test" {
		t.Errorf("editCommitMessage is not test %s", res)
	}
}
