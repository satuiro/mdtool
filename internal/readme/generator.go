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

func (g *Generator) shouldIncludeFile(path string, size int) bool { // Changed from int64 to int
	if size > g.maxFileSize {
		return false
	}

	for _, pattern := range g.excludePatterns {
		if strings.HasSuffix(path, pattern) || strings.Contains(path, pattern) {
			return false
		}
	}

	return true
}

func (g *Generator) getRepoFiles(ctx context.Context) (map[string]string, *github.Repository, error) {
	parts := strings.Split(g.repoName, "/")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid repository name format. Use 'owner/repo'")
	}

	repo, _, err := g.githubClient.Repositories.Get(ctx, parts[0], parts[1])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get repository: %w", err)
	}

	files := make(map[string]string)
	_, dirContent, _, err := g.githubClient.Repositories.GetContents(ctx, parts[0], parts[1], "", nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get repository contents: %w", err)
	}

	for _, content := range dirContent {
		if content.GetType() == "file" {
			if !g.shouldIncludeFile(content.GetPath(), content.GetSize()) {
				continue
			}

			fileContent, _, _, err := g.githubClient.Repositories.GetContents(ctx, parts[0], parts[1], content.GetPath(), nil)
			if err != nil {
				color.Yellow("WARNING: Skipping file %s: %v", content.GetPath(), err)
				continue
			}

			decoded, err := base64.StdEncoding.DecodeString(*fileContent.Content)
			if err != nil {
				color.Yellow("WARNING: Failed to decode file %s: %v", content.GetPath(), err)
				continue
			}

			files[content.GetPath()] = string(decoded)
			color.Blue("INFO: Added file: %s", content.GetPath())
		}
	}

	return files, repo, nil
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
