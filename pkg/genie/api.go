package genie

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

var client *resty.Request

var apiURL string = os.Getenv("OPENAI_HOST")
var apiToken string = os.Getenv("OPENAI_API_KEY")

var commitMessages []string

func InitClient() *resty.Request {
	client = resty.New().SetBaseURL("https://"+apiURL).R().
		EnableTrace().
		SetAuthScheme("Bearer").
		SetAuthToken(apiToken).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("User-Agent", "git-genie/0.0.1")
	return client
}

func SubmitToApi(diff string) []string {
	// DBEUG
	//var prompt string = "You are a programmer and want to commit this code. Write a commit message for these code changes. \nCode changes: \n--- a/client/src/components/apps/new.vue\n+++ b/client/src/components/apps/new.vue\n@@ -841,13 +841,14 @@ export default {\n */\n     }),\n     mounted() {\n-      if (this.$route.query.service) {\n-        this.loadTemplate(this.$route.query.service);\n-      }\n       this.loadStorageClasses();\n       this.loadPipeline();\n       this.loadPodsizeList();\n       this.loadApp(); // this may lead into a race condition with the buildpacks loaded in loadPipeline\n+\n+      if (this.$route.query.service) {\n+        this.loadTemplate(this.$route.query.service);\n+      }\n     },\n     components: {\n         Addons,\n@@ -868,6 +869,14 @@ export default {\n           this.cronjobs = response.data.cronjobs;\n           this.addons = response.data.addons;\n\n+          if (response.data.image.build) {\n+            console.log(\"buildpack build\", response.data.image.build);\n+            this.buildpack.build = response.data.image.build;\n+          }\n+\n+          if (response.data.image.run) {\n+            this.buildpack.run = response.data.image.run;\n+          }\n\n           // Open Panel if there is some data to show\n           if (this.envvars.length > 0) {"
	//var prompt string = "Say this is a test"
	//diff = "--- a/client/src/components/apps/new.vue\n+++ b/client/src/components/apps/new.vue\n@@ -841,13 +841,14 @@ export default {\n */\n     }),\n     mounted() {\n-      if (this.$route.query.service) {\n-        this.loadTemplate(this.$route.query.service);\n-      }\n       this.loadStorageClasses();\n       this.loadPipeline();\n       this.loadPodsizeList();\n       this.loadApp(); // this may lead into a race condition with the buildpacks loaded in loadPipeline\n+\n+      if (this.$route.query.service) {\n+        this.loadTemplate(this.$route.query.service);\n+      }\n     },\n     components: {\n         Addons,\n@@ -868,6 +869,14 @@ export default {\n           this.cronjobs = response.data.cronjobs;\n           this.addons = response.data.addons;\n\n+          if (response.data.image.build) {\n+            console.log(\"buildpack build\", response.data.image.build);\n+            this.buildpack.build = response.data.image.build;\n+          }\n+\n+          if (response.data.image.run) {\n+            this.buildpack.run = response.data.image.run;\n+          }\n\n           // Open Panel if there is some data to show\n           if (this.envvars.length > 0) {"
	//var prompt string = "You are a programmer and want to commit this code. Write a commit message for these code changes.\n\n" + diff
	var prompt string = "You are a programmer and want to commit this code. Describe the code changes in one sentence.\n\n" + diff
	var jsonPrompt string = jsonEscape(prompt)
	var CompletionResponse CompletionResponse
	var body = `{
		"model": "text-davinci-003",
		"prompt": "` + jsonPrompt + `",
		"temperature": 1,
		"max_tokens": 300,
		"top_p": 1,
		"n": 3,
		"stream": false,
		"echo": false,
		"frequency_penalty": 0,
		"presence_penalty": 0
	  }
	  `
	client.SetBody(body)
	client.SetResult(&CompletionResponse)

	resp, err := client.Post("/v1/completions")
	if err != nil {
		panic(err)
	}
	println(resp.Status())

	// DEBUG
	//b := resp.Body()
	//println(string(b))

	for _, choice := range CompletionResponse.Choices {
		//println(choice.Text)

		commitMessages = append(commitMessages, strings.TrimLeft(choice.Text, "\n"))
	}

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
