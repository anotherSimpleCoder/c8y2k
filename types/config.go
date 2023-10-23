package types

import (
	"encoding/json"
	"os"
)

type Config struct {
	ProjectName string
	ProjectType string

	TenantURL string
	Username  string
	Password  string
	TenantID  string

	Path    string
	Modules []string
}

func ReadConfig() *Config {
	var cfg Config

	if ConfigExists() {
		content, err := os.ReadFile("c8y2k.json")
		if err != nil {
			return nil
		}

		if err := json.Unmarshal(content, &cfg); err != nil {
			return nil
		}
	}

	return &cfg
}

func ConfigExists() bool {
	if _, err := os.OpenFile("c8y2k.json", 0755, os.ModeAppend); err != nil {
		return false
	}

	return true
}
