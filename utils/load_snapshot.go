package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func LoadSnapshot(commitID string) (map[string]string, error) {
	snapshot := make(map[string]string)

	if commitID == "" {
		return nil, nil
	}

	path := filepath.Join(".litevc", "commits", commitID) // mengambil snapshot dari commit yang sudah ada di .litevc/commits/{commitID}

	data, err := os.ReadFile(path) // membaca file snapshot

	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}

		snapshot[parts[0]] = parts[1]
	}

	return snapshot, nil
}
