<p align="center"><img src="docs/img/mascot.png" height="200" /></p>

# GitGenie
Once upon a time, there was a magical genie who lived inside Git as a plugin. This genie loved to help developers write better commit messages when they changed their code.

The genie was very clever and used its magic ChatGPT powers to analyze the changes in the code and suggest good messages for the developers to use.

And if the developers weren't happy with any of the suggestions, he would open an editor and give them the opportunity to adjust the suggestions to their liking.

> **Note**
> The genie needs your help: There are binaries for Linux and Windows (ARM and x86_64). I have only a Mac(x86_64) to test the binaries. If you have a Windows/Linux or ARM architecture [please let me know in the issues section](https://github.com/kubero-dev/GitGenie/issues/2) if it works.

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
<img src="docs/img/demo.gif" />

## Installation

Download the binary for you operating system from the [releases page](https://github.com/mms-gianni/GitGenie/releases/latest) and extract it to /usr/local/bin or any other directory in your PATH.

```bash 
tar -xvzf GitGenie_Mac_x86_64.tar.gz -C /usr/local/bin

export OPENAI_API_KEY=sk-y..............................
```

Add the OpenAI Key ENV to your .bashrc or .zshrc and configure your Genie with the following additional ENV variables to change the default behavior.

- `OPENAI_API_KEY`: OpenAI API token **(required)**
- `OPENAI_HOST`: OpenAI API host (default: `api.openai.com`)
- `EDITOR` or `VISUAL`: Editor to edit commit message (default: `vim`)
- `GENIE_SUGESTIONS`: Number of suggestions to generate (default: `3`)
- `GENIE_LENGTH`: Length of each suggestion (default: `medium`, can be `short`, `medium`, `long`, `verylong`)
- `GENIE_MAX_TOKENS`: Maximum number of tokens to generate (overrides `GENIE_LENGTH`)
- `GENIE_SKIP_EDIT`: Skip editing the commit message (default: `false`)
- `GENIE_LANGUAGE`: Language to use for the commit message (default: `en`, can be `en`, `ch`, `de`, `es`, `fr`, `it`, `ja`, `ko`, `pt`, `ru`, `zh`)

## Privacy
There may be many reasons a repository owner does not want to allow genie generated commit messages or code to be sent to a third party service. For this reason, the repository can be configured to block GitGenie completly.

To block the genie, create a file called `.gitgenieblock` in the root of your repository and add the following content:

```bash
touch .gitgenieblock
```

## Repository Configuration (optional)
Prepare your repository for GitGenie by adding a `.gitgenie` file to the root of your repository. This file can be used to configure the genie for your repository.

Parameters: 
- `language`: Language to use for the commit message (default: `en`, can be `en`, `ch`, `de`, `es`, `fr`, `it`, `ja`, `ko`, `pt`, `ru`, `zh`)
- `description`: Description of the repository/software (default: `""`)

## Examples

Create a commit message with 5 suggestions and skip editing the commit message
```bash
git gci -n 5 -f
```

Create a commit message in a different language and use a short commit message
```bash
git gci -L de -l short
```

