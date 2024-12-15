package readme

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/v60/github"
	"github.com/satuiro/mdtool/internal/config"
	"github.com/satuiro/mdtool/internal/groq"
	"golang.org/x/oauth2"
)

type Generator struct {
	config          *config.Config
	githubClient    *github.Client
	repoName        string
	maxFileSize     int // Changed from int64 to int
	excludePatterns []string
}

func NewGenerator(config *config.Config, repoName string) *Generator {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &Generator{
		config:       config,
		githubClient: github.NewClient(tc),
		repoName:     repoName,
		maxFileSize:  100000, // 100KB
		excludePatterns: []string{
			"node_modules/",
			"venv/",
			".git/",
			"__pycache__/",
			".pyc",
			".pyo",
			".pyd",
			".so",
			".dylib",
			".dll",
		},
	}
}

func (g *Generator) getRepoFiles(ctx context.Context) (map[string]string, *github.Repository, error) {
	parts := strings.Split(g.repoName, "/")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid repository name format. Use 'owner/repo'")
	}

	owner, repoName := parts[0], parts[1]

	// Get repository information
	repo, _, err := g.githubClient.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get repository: %w", err)
	}

	files := make(map[string]string)

	// Create a recursive function to traverse directories
	var fetchContents func(path string) error
	fetchContents = func(path string) error {
		_, dirContent, _, err := g.githubClient.Repositories.GetContents(ctx, owner, repoName, path, nil)
		if err != nil {
			return fmt.Errorf("failed to get contents for path %s: %w", path, err)
		}

		for _, content := range dirContent {
			if !g.shouldIncludeFile(content.GetPath(), int64(content.GetSize())) {
				color.Yellow("WARNING: Skipping file %s (size or pattern excluded)", content.GetPath())
				continue
			}

			switch content.GetType() {
			case "file":
				fileContent, _, _, err := g.githubClient.Repositories.GetContents(ctx, owner, repoName, content.GetPath(), nil)
				if err != nil {
					color.Yellow("WARNING: Skipping file %s: %v", content.GetPath(), err)
					continue
				}

				if fileContent.GetEncoding() == "base64" {
					decoded, err := base64.StdEncoding.DecodeString(*fileContent.Content)
					if err != nil {
						color.Yellow("WARNING: Failed to decode file %s: %v", content.GetPath(), err)
						continue
					}
					files[content.GetPath()] = string(decoded)
					color.Blue("INFO: Added file: %s", content.GetPath())
				}

			case "dir":
				// Recursively fetch contents of subdirectory
				if err := fetchContents(content.GetPath()); err != nil {
					color.Yellow("WARNING: Failed to fetch contents of directory %s: %v", content.GetPath(), err)
				}
			}
		}
		return nil
	}

	// Start recursive fetch from root
	if err := fetchContents(""); err != nil {
		return nil, nil, fmt.Errorf("failed to fetch repository contents: %w", err)
	}

	if len(files) == 0 {
		return nil, nil, fmt.Errorf("no suitable files found in the repository")
	}

	color.Green("INFO: Successfully fetched %d files from repository", len(files))
	return files, repo, nil
}

// shouldIncludeFile checks if a file should be included based on size and path
func (g *Generator) shouldIncludeFile(path string, size int64) bool {
	// Increase max file size to 500KB
	if size > 500000 {
		return false
	}

	// Common binary and generated file patterns to exclude
	excludePatterns := []string{
		"node_modules/",
		"venv/",
		".git/",
		"__pycache__/",
		"target/",
		"dist/",
		"build/",
		".idea/",
		".vscode/",
		".DS_Store",
		"Thumbs.db",
		".env",
		".pyc",
		".pyo",
		".pyd",
		".so",
		".dylib",
		".dll",
		".exe",
		".bin",
		".dat",
		".pb",
		".o",
		".a",
		".lib",
		".png",
		".jpg",
		".jpeg",
		".gif",
		".ico",
		".svg",
		".woff",
		".woff2",
		".ttf",
		".eot",
	}

	// Check if the file matches any exclude pattern
	for _, pattern := range excludePatterns {
		if strings.Contains(strings.ToLower(path), strings.ToLower(pattern)) {
			return false
		}
	}

	return true
}

func (g *Generator) Generate() (string, error) {
	ctx := context.Background()

	files, repo, err := g.getRepoFiles(ctx)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		color.Yellow("WARNING: No suitable files found for analysis")
		return "", nil
	}

	groqService := groq.NewService(g.config.GroqAPIKey)
	content, err := groqService.GenerateReadme(files, repo)
	if err != nil {
		return "", fmt.Errorf("failed to generate README: %w", err)
	}

	return content, nil
}

func (g *Generator) DisplayReadme(content, format string) {
	if format == "raw" {
		fmt.Println(content)
	} else {
		// Add some basic formatting
		fmt.Printf("\n%s\n\n%s\n", color.GreenString("Generated README:"), content)
	}
}
