<p align="center"><img src="docs/img/mascot.png" height="200" /></p>

# GitGenie
GitGenie is a git plugin that creates commit messages with ChatGPT. It is a replacement command for `git commit`. The shorthand gci stands for "generate commit" 

Not happy with the suggested commit message? No problem, GitGenie will open your editor and let you edit the commit message or leave it empty.

## Usage
    
```bash
git add . 
git gci 
```

```
git gci [flags]

Flags:
  -f, --fast                 Skip editing the commit message
  -h, --help                 help for git
  -H, --host string          OpenAI API host (default "api.openai.com")
  -L, --language string      Commit message language: en, ch, de, es, fr, it, ja, ko, pt, zh (default "en")
  -l, --length string        Commit message length: short, medium, long, verylong (default "medium")
  -s, --signoff              Add signing signature to commit message
  -n, --suggestions string   Number of suggestions to generate (default "3")
```


## Installation

Download the Archived binary for you operating system from the [releases page](https://github.com/mms-gianni/GitGenie/releases/latest) and extract it to /usr/local/bin or any other directory in your PATH.

```bash 
tar -xvzf GitGenie_Mac_x86_64.tar.gz -C /usr/local/bin

export OPENAI_API_KEY=sk-y..............................
```

Add the OpenAI Key ENV your .bashrc or .zshrc and configure your Genie with the following additional ENV variables to change the default behavior.

- `OPENAI_API_KEY`: OpenAI API token **(required)**
- `OPENAI_HOST`: OpenAI API host (default: `api.openai.com`)
- `EDITOR` or `VISUAL`: Editor to edit commit message (default: `vim`)
- `GENIE_SUGESTIONS`: Number of suggestions to generate (default: `3`)
- `GENIE_LENGTH`: Length of each suggestion (default: `medium`, can be `short`, `medium`, `long`, `verylong`)
- `GENIE_MAX_TOKENS`: Maximum number of tokens to generate (overrides `GENIE_LENGTH`)
- `GENIE_SKIP_EDIT`: Skip editing the commit message (default: `false`)
- `GENIE_LANGUAGE`: Language to use for the commit message (default: `en`, can be `en`, `ch`, `de`, `es`, `fr`, `it`, `ja`, `ko`, `pt`, `ru`, `zh`)


## Examples

Create a commit message with 5 suggestions and skip editing the commit message
```
git gci -s 5 -e
```

Create a commit message in a different language and use a short commit message
```
git gci -L de -l short
```