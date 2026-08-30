package main

// Answers: the four questions, each 0..2, saved between runs.

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func defaultPath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "answers.json"
	}
	return filepath.Join(dir, "kindred", "answers.json")
}

func LoadAnswers() (Answers, error) {
	var a Answers
	data, err := os.ReadFile(defaultPath())
	if err != nil {
		return a, err
	}
	err = json.Unmarshal(data, &a)
	return a, err
}

func SaveAnswers(a Answers) error {
	path := defaultPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, _ := json.MarshalIndent(a, "", "  ")
	return os.WriteFile(path, data, 0o644)
}
