package config

import "os"

type Config struct {
	GroqAPIKey   string
	GithubToken  string
	DefaultModel string
}

func GetConfig() *Config {
	return &Config{
		GroqAPIKey:   getEnvOrDefault("GROQ_API_KEY", ""),
		GithubToken:  getEnvOrDefault("GITHUB_TOKEN", ""),
		DefaultModel: "llama-3.2-90b-vision-preview",
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
