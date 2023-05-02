# GitGenie
Create commit messages with ChatGPT

Usage:
    
```bash
git gci 
```

Available env variables:

- `OPENAI_API_KEY`: OpenAI API token **(required)**
- `OPENAI_HOST`: OpenAI API host (default: `api.openai.com`)
- `EDITOR` or `VISUAL`: Editor to edit commit message (default: `vim`)
- `GENIE_SUGESTIONS`: Number of suggestions to generate (default: `3`)
- `GENIE_LENGTH`: Length of each suggestion (default: `medium`, can be `short`, `medium`, `long`, `verylong`)

installation instructions:

Download the Archived binary from the [releases page]() and extract it to /usr/local/bin

```bash 
sudo tar -xvf git-genie.tar.gz -C /usr/local/bin

export OPENAI_API_KEY=sk-y..............................
```