package utils

import (
	"os"
	"log"
	"path/filepath"
)

func WriteFile(path string, data []byte) {
	log.Printf("Writing to %s", path)
	os.MkdirAll(filepath.Dir(path), 0755)

	os.WriteFile(path, []byte(data), 0644)
}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}