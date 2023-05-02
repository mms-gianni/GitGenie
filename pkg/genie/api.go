package genie

import (
	"encoding/json"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/leaanthony/spinner"
)

var client *resty.Request

var commitMessages []string

func InitClient() *resty.Request {
	loadConfig()

	client = resty.New().SetBaseURL("https://"+config.openAiApiHost).R().
		EnableTrace().
		SetAuthScheme("Bearer").
		SetAuthToken(config.openAiApiToken).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", "git-genie/0.0.1")
	return client
}

func SubmitToApi(diff string) []string {

	var prompt string = "You are a programmer and want to commit this code. Describe the code changes in one sentence.\n\n" + diff
	var jsonPrompt string = jsonEscape(prompt)
	var CompletionResponse CompletionResponse
	var body = `{
		"model": "text-davinci-003",
		"prompt": "` + jsonPrompt + `",
		"temperature": 1,
		"max_tokens": 300,
		"top_p": 1,
		"n": ` + config.suggestions + `,
		"stream": false,
		"echo": false,
		"frequency_penalty": 0,
		"presence_penalty": 0
	  }
	  `
	client.SetBody(body)
	client.SetResult(&CompletionResponse)

	s := spinner.New()
	s.Start("Loading commit messages...")
	resp, err := client.Post("/v1/completions")
	if err != nil {
		s.Error("Error loading commit messages[" + resp.Status() + "]")
		panic(err)
	}
	s.Success("Commit messages loaded [" + resp.Status() + "]")

	for _, choice := range CompletionResponse.Choices {
		commitMessages = append(commitMessages, strings.TrimLeft(choice.Text, "\n"))
	}
	commitMessages = append(commitMessages, "<empty>")

	return commitMessages
}

type CompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
}
