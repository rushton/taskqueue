package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func New() (string, error) {
	var dir string
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	switch runtime.GOOS {
	case "windows":
		dir = "AppData/Roaming"
	case "linux":
		dir = ".local/share"
	default:
		return "", fmt.Errorf("unsupported operating system %s", runtime.GOOS)
	}
	return filepath.Join(home, dir), nil
}
