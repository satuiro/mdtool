package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/v60/github"
)

type Service struct {
	apiKey string
}

func NewService(apiKey string) *Service {
	return &Service{
		apiKey: apiKey,
	}
}

type GroqRequest struct {
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}

func (s *Service) GenerateReadme(files map[string]string, repo *github.Repository) (string, error) {
	// Validate API key
	if s.apiKey == "" {
		return "", fmt.Errorf("GROQ_API_KEY is not set")
	}

	// Build file summaries
	var fileSummaries []string
	for path, content := range files {
		preview := content
		if len(content) > 200 {
			preview = content[:200] + "..."
		}
		fileSummaries = append(fileSummaries, fmt.Sprintf("### %s\n```\n%s\n```", path, preview))
	}

	// Build metadata
	metadata := []string{
		fmt.Sprintf("- **name**: %s", repo.GetName()),
		fmt.Sprintf("- **description**: %s", repo.GetDescription()),
		fmt.Sprintf("- **language**: %s", repo.GetLanguage()),
		fmt.Sprintf("- **license**: %s", func() string {
			if repo.License != nil {
				return repo.License.GetName()
			}
			return "None"
		}()),
		fmt.Sprintf("- **stars**: %d", repo.GetStargazersCount()),
		fmt.Sprintf("- **forks**: %d", repo.GetForksCount()),
		fmt.Sprintf("- **open_issues**: %d", repo.GetOpenIssuesCount()),
	}

	prompt := fmt.Sprintf(`Generate a comprehensive README.md for the following project:
Repository Metadata:
%s

Project Files:
%s

Create a detailed README that includes:
1. Project title and description
2. Key features
3. Installation instructions
4. Usage examples
5. Project structure overview
6. Dependencies and requirements
7. Contributing guidelines (if applicable)
8. License information (if available)

Use clear markdown formatting with appropriate sections. Focus on creating a helpful
and informative README that would help users understand and use the project.
If the project has a specific focus or unique features, highlight those prominently.`,
		strings.Join(metadata, "\n"),
		strings.Join(fileSummaries, "\n\n"))

	color.Blue("INFO: Sending request to Groq API...")

	reqBody := GroqRequest{
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Model:       "llama-3.3-70b-versatile", // Updated to use a more reliable model
		Temperature: 0.7,
		MaxTokens:   2000,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Log response status
	color.Blue("INFO: Groq API response status: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var groqResp GroqResponse
	if err := json.Unmarshal(bodyBytes, &groqResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	// Check for API error response
	if groqResp.Error.Message != "" {
		return "", fmt.Errorf("API error: %s (type: %s)", groqResp.Error.Message, groqResp.Error.Type)
	}

	if len(groqResp.Choices) == 0 {
		return "", fmt.Errorf("API returned no choices in response: %s", string(bodyBytes))
	}

	content := groqResp.Choices[0].Message.Content
	if content == "" {
		return "", fmt.Errorf("API returned empty content: %s", string(bodyBytes))
	}

	color.Green("INFO: Successfully generated README content")
	return content, nil
}
