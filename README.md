<p align="center"><img src="docs/img/mascot.png" height="200" /></p>

# GitGenie
GitGenie is a git plugin that creates commit messages with ChatGPT. It is a replacement command for `git commit`. The shorthand gci stands for "generate commit" 

## Usage
    
```bash
git add . 
git gci 
```

## Installation instructions

Download the Archived binary from the [releases page](https://github.com/mms-gianni/GitGenie/releases/latest) and extract it to /usr/local/bin or any other directory in your PATH.

```bash 
tar -xvzf GitGenie_Mac_x86_64.tar.gz -C /usr/local/bin

export OPENAI_API_KEY=sk-y..............................
```

Configure GitGenie in your .bashrc or .zshrc with the following env variables:


### Configuration

Available env variables:

- `OPENAI_API_KEY`: OpenAI API token **(required)**
- `OPENAI_HOST`: OpenAI API host (default: `api.openai.com`)
- `EDITOR` or `VISUAL`: Editor to edit commit message (default: `vim`)
- `GENIE_SUGESTIONS`: Number of suggestions to generate (default: `3`)
- `GENIE_LENGTH`: Length of each suggestion (default: `medium`, can be `short`, `medium`, `long`, `verylong`)
- `GENIE_MAX_TOKENS`: Maximum number of tokens to generate (overrides `GENIE_LENGTH`)
- `GENIE_SKIP_EDIT`: Skip editing the commit message (default: `false`)
