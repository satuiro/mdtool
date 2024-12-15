# mdtool - GitHub README Generator

A CLI tool to automatically generate comprehensive README files for GitHub repositories.

## Installation

### Using `go install`

```bash
go install github.com/satuiro/mdtool/cmd/mdtool@latest
```

### Building from source

1. Clone the repository:

   ```bash
   git clone https://github.com/satuiro/mdtool.git
   cd mdtool
   ```

2. Build and install:
   ```bash
   make install
   ```

## Usage

1. Set required environment variables:

   ```bash
   export GROQ_API_KEY="your-groq-api-key"
   export GITHUB_TOKEN="your-github-token"
   ```

2. Generate README for a repository:
   ```bash
   mdtool readme -r owner/repo
   ```

### Options

- `-r, --repo`: Repository name in format 'owner/repo' (required)
- `-o, --output`: Output format (preview/raw)
- `-s, --save`: Save README to file
- `--version`: Show version information

