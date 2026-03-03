package utils

import (
	"fmt"
	datastructures "litevc/data_structures"
	"os"
	"strings"
)

func WriteIndex(entries []datastructures.IndexEntry) error {
	if len(entries) <= 0 {
		fmt.Println("No any changes")
		return nil
	}

	// Here we can write the index entries to a file index
	var lines []string

	for _, entry := range entries {
		if entry.Action == "DELETED" {
			lines = append(lines, fmt.Sprintf("%s %s", entry.Action, entry.FilePath))
		} else {
			lines = append(lines, fmt.Sprintf("%s %s %s", entry.Action, entry.FilePath, entry.HashedFile))
		}
	}

	content := strings.Join(lines, "\n")

	return os.WriteFile(".litevc/index", []byte(content), 0644)
}
