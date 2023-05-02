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

installation instructions:

```bash
curl https://raw.githubusercontent.com/roerohan/gitgenie/master/install.sh | bash

export OPENAI_API_KEY=sk-y..............................
```