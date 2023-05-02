# GitGenie
Create commit messages with ChatGPT

## Usage
    
```bash
git add . 
git gci 
```

## Installation instructions

Download the Archived binary from the [releases page](https://github.com/mms-gianni/GitGenie/releases) and extract it to /usr/local/bin

```bash 
sudo tar -xvf git-genie.tar.gz -C /usr/local/bin

export OPENAI_API_KEY=sk-y..............................
```

## Configuration

Available env variables:

- `OPENAI_API_KEY`: OpenAI API token **(required)**
- `OPENAI_HOST`: OpenAI API host (default: `api.openai.com`)
- `EDITOR` or `VISUAL`: Editor to edit commit message (default: `vim`)
- `GENIE_SUGESTIONS`: Number of suggestions to generate (default: `3`)
- `GENIE_LENGTH`: Length of each suggestion (default: `medium`, can be `short`, `medium`, `long`, `verylong`)
- `GENIE_MAX_TOKENS`: Maximum number of tokens to generate (overrides `GENIE_LENGTH`)
- `GENIE_SKIP_EDIT`: Skip editing the commit message (default: `false`)