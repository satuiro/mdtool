# CLI Tool: mdtool

## Core Tech Stack

1. Go (core language)
2. Libraries:
   - github.com/charmbracelet/glamour - Markdown rendering with styles
   - github.com/charmbracelet/lipgloss - Terminal styling
   - github.com/spf13/cobra - CLI framework
   - github.com/fatih/color - Color output
   - github.com/yuin/goldmark - Markdown parsing
   - github.com/go-git/go-git/v5 - Git operations
   - github.com/fsnotify/fsnotify - File watching

## Features (In Order of Implementation)

### Phase 1: Basic Markdown Preview

- [x] Parse markdown file
- [ ] Render headers with proper styling
- [ ] Handle basic formatting (bold, italic, strikethrough)
- [ ] Support code blocks with syntax highlighting
- [ ] Display lists (ordered and unordered)
- [ ] Show tables with alignment
- [ ] Handle links (with color differentiation)
- [ ] Support blockquotes

### Phase 2: Git Integration

- [ ] Detect git repository
- [ ] Read repository structure
- [ ] Parse all code files
- [ ] Extract relevant information for README
- [ ] Analyze project structure
- [ ] Identify main technologies used

### Phase 3: README Generation

- [ ] Connect to Groq API
- [ ] Generate project description
- [ ] Create installation instructions
- [ ] Add usage examples
- [ ] List dependencies
- [ ] Generate contributing guidelines
- [ ] Add license information

### Phase 4: Real-time Features

- [ ] Watch file for changes
- [ ] Live preview updates
- [ ] Cache rendered content
- [ ] Handle partial updates

### Phase 5: Advanced Features

- [ ] Custom themes support
- [ ] Export to different formats
- [ ] Table of contents generation
- [ ] Dead link checking
- [ ] Spelling/grammar suggestions
- [ ] Markdown linting

## Initial Project Structure

```
.
├── cmd/
│   └── mdtool/
│       └── main.go           # Entry point
├── internal/
│   ├── render/
│   │   └── render.go         # Markdown rendering logic
│   ├── parser/
│   │   └── parser.go         # Markdown parsing
│   ├── git/
│   │   └── git.go           # Git operations
│   ├── groq/
│   │   └── client.go        # Groq API client
│   ├── style/
│   │   └── style.go         # Terminal styling
│   └── config/
│       └── config.go        # Configuration handling
├── pkg/
│   ├── markdown/
│   │   └── markdown.go      # Public markdown utilities
│   └── utils/
│       └── utils.go         # Shared utilities
└── go.mod
```

## Basic Command Structure

```go
// Example commands
mdtool preview README.md              // Preview markdown
mdtool preview --watch README.md      // Watch mode
mdtool generate --repo .              // Generate README
mdtool style --theme dark README.md   // Custom styling

```
