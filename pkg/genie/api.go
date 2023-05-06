package genie

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/leaanthony/spinner"
	log "github.com/pieterclaerhout/go-log"
)

var client *resty.Request

var commitMessages []string

func initClient() *resty.Request {

	client = resty.New().SetBaseURL("https://"+config.OpenAiApiHost).R().
		EnableTrace().
		SetAuthScheme("Bearer").
		SetAuthToken(config.OpenAiApiToken).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", "git-genie/0.0.1")
	return client
}

func submitToApiChat(diff string) []string {
	var gitRoot = getGitRoot()
	var repoConfig repoConfig
	repoConfig.loadRepoConfig(gitRoot)

	// check if file exists
	if _, err := os.Stat(gitRoot + "/.gitgenieblock"); err == nil {
		fmt.Println(gitRoot + "/.gitgenieblock")
		fmt.Println("This repository does not allow genie commits.")
		os.Exit(1)
	}

	if repoConfig.Language != "" {
		config.Language = repoConfig.Language
	}

	var prompt string = getUser(config.Language) + "\n\n" + diff
	var system string = getSystem(config.Language)
	var jsonPrompt string = jsonEscape(prompt)
	var ChatCompletionResponse ChatCompletionResponse

	if value, ok := os.LookupEnv("GENIE_MAX_TOKENS"); ok {
		config.Max_tokens = value
	} else {
		switch config.Length {
		case "veryshort", "vs":
			config.Max_tokens = "50"
		case "short", "s":
			config.Max_tokens = "100"
		case "medium", "m":
			config.Max_tokens = "300"
		case "long", "l":
			config.Max_tokens = "500"
		case "verylong", "vl":
			config.Max_tokens = "1000"
		default:
			config.Max_tokens = "301"
		}
	}

	var messages = []Message{
		{
			Role:    "system",
			Content: system,
		},
		{
			Role:    "system",
			Content: repoConfig.Description,
		},
	}

	var commitLog = getLogs()
	for _, entry := range commitLog {
		messages = append(messages, Message{
			Role:    "assistant",
			Content: entry,
		})
	}

	messages = append(messages, Message{
		Role:    "user",
		Content: jsonPrompt,
	})

	m, _ := json.Marshal(messages)

	var body = `{
		"model": "gpt-3.5-turbo",
		"messages": ` + string(m) + `,
		"temperature": 1,
		"max_tokens": ` + config.Max_tokens + `,
		"top_p": 1,
		"n": ` + config.Suggestions + `,
		"stream": false,
		"frequency_penalty": 0,
		"presence_penalty": 0
	  }
	`
	log.DebugDump("Request Body", body)

	client.SetBody(body)
	client.SetResult(&ChatCompletionResponse)

	s := spinner.New()
	s.Start("Loading commit messages...")
	resp, err := client.Post("/v1/chat/completions")
	if err != nil {
		s.Error("Error loading commit messages[" + resp.Status() + "]")
		panic(err)
	}
	if resp.StatusCode() > 299 {
		s.Error("Error loading commit messages[" + resp.Status() + "]")
		fmt.Println(body)
		panic(string(resp.Body()))
	}
	log.DebugDump("Response Body", string(resp.Body()))
	s.Success("Commit messages loaded [" + resp.Status() + "]")

	for _, choice := range ChatCompletionResponse.Choices {
		commitMessages = append(commitMessages, choice.Message.Content+"\n")
	}
	commitMessages = append(commitMessages, "<empty>")

	return commitMessages
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int                `json:"created"`
	Model   string             `json:"model"`
	Choices []CompletionChoice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type CompletionChoice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type ChatCompletionResponse struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []ChatChoice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type ChatChoice struct {
	Index   int `json:"index"`
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	FinishReason string `json:"finish_reason"`
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
}
