package genie

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/leaanthony/spinner"
)

var client *resty.Request

var commitMessages []string

func initClient() *resty.Request {

	client = resty.New().SetBaseURL("https://"+config.openAiApiHost).R().
		EnableTrace().
		SetAuthScheme("Bearer").
		SetAuthToken(config.openAiApiToken).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", "git-genie/0.0.1")
	return client
}

func submitToApi(diff string) []string {

	var prompt string = "You are a programmer and want to commit this code. Describe the code changes in one sentence.\n\n" + diff
	var jsonPrompt string = jsonEscape(prompt)
	var CompletionResponse CompletionResponse
	var body = `{
		"model": "text-davinci-003",
		"prompt": "` + jsonPrompt + `",
		"temperature": 1,
		"max_tokens": ` + config.max_tokens + `,
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

func submitToApiChat(diff string) []string {
	var prompt string = getUser(config.language) + "\n\n" + diff
	var system string = getSystem(config.language)
	var jsonPrompt string = jsonEscape(prompt)
	var ChatCompletionResponse ChatCompletionResponse
	var body = `{
		"model": "gpt-3.5-turbo",
		"messages": [
			{
				"role": "system", 
				"content": "` + system + `"
			},
			{
				"role": "user", 
				"content": "` + jsonPrompt + `"
			}
		],
		"temperature": 1,
		"max_tokens": ` + config.max_tokens + `,
		"top_p": 1,
		"n": ` + config.suggestions + `,
		"stream": false,
		"frequency_penalty": 0,
		"presence_penalty": 0
	  }
	  `
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
	s.Success("Commit messages loaded [" + resp.Status() + "]")

	for _, choice := range ChatCompletionResponse.Choices {
		commitMessages = append(commitMessages, strings.TrimLeft(choice.Message.Content, "\n"))
	}
	commitMessages = append(commitMessages, "<empty>")

	return commitMessages
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
