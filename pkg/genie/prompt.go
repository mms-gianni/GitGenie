package genie

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed language.yaml
var languagesFile string

var languages []Language

type Language struct {
	Language string `yaml:"language"`
	System   string `yaml:"system"`
	User     string `yaml:"user"`
}

func loadFromLanguageYaml() {
	// Load the language yaml file
	// unmarshal the yaml file into a struct
	yaml.Unmarshal([]byte(languagesFile), &languages)
	//fmt.Println(getLanguagesList())

}

/*
func getLanguagesList() []string {
	var languageList []string
	for _, language := range languages {
		languageList = append(languageList, language.Language)
	}
	return languageList
}
*/

func getSystem(language string) string {
	for _, lang := range languages {
		if lang.Language == language {
			return lang.System
		}
	}
	return ""
}

func getUser(language string) string {
	for _, lang := range languages {
		if lang.Language == language {
			return lang.User
		}
	}
	return ""
}
