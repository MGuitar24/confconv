package confconv

import (
	"os"
	"path/filepath"
	"strings"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func identifyFileType(path string) string {
	if fileExists(path) {
		return strings.ToLower(filepath.Ext(path)[1:])
	}
	return "unknown"
}

func readFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	return file, err
}

func writeFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
