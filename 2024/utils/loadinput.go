package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadInput(file string) (string, error) {
	// todo make year agnostic
	// wd, _ := os.Getwd()
	// log.Printf("Current working directory: %s\n", wd)

	filePath := filepath.Join("2024", "input", file)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to load input for %s: %w", file, err)
	}

	return string(data), nil
}
