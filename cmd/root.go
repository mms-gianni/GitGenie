package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git gci",
	Short: "GitGenie is a git plugin that creates commit messages with ChatGPT.",
	Long: `Improve your commit messages with GitGenie. 
	
GitGenie is a git plugin that creates commit messages with ChatGPT.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Suggestions string
var Length string
var Signoff bool
var Fast bool
var OpenAiApiHost string
var OpenAiApiToken string
var MaxTokens string
var Language string

func init() {
	OpenAiApiHost = getEnv("OPENAI_API_HOST", "api.openai.com")
	rootCmd.Flags().StringVarP(&OpenAiApiHost, "host", "H", OpenAiApiHost, "OpenAI API host")

	OpenAiApiToken = os.Getenv("OPENAI_API_KEY")

	rootCmd.Flags().BoolVarP(&Signoff, "signoff", "s", false, "Add signing signature to commit message")

	Suggestions = getEnv("GENIE_SUGESTIONS", "3")
	rootCmd.Flags().StringVarP(&Suggestions, "suggestions", "n", Suggestions, "Number of suggestions to generate")

	Length = getEnv("GENIE_LENGTH", "medium")
	rootCmd.Flags().StringVarP(&Length, "length", "l", Length, "Commit message length: short, medium, long, verylong")

	if os.Getenv("GENIE_SKIP_EDIT") == "true" {
		Fast = true
	}
	rootCmd.Flags().BoolVarP(&Fast, "fast", "f", Fast, "Skip editing the commit message")

	if value, ok := os.LookupEnv("GENIE_MAX_TOKENS"); ok {
		MaxTokens = value
	} else {
		switch Length {
		case "short":
			MaxTokens = "100"
		case "medium":
			MaxTokens = "300"
		case "long":
			MaxTokens = "500"
		case "verylong":
			MaxTokens = "1000"
		default:
			MaxTokens = "300"
		}
	}

	Language = getEnv("GENIE_LANGUAGE", "en")
	rootCmd.Flags().StringVarP(&Language, "language", "L", Language, "Commit message language: en, ch, de, es, fr, it, ja, ko, pt, zh")
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
