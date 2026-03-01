package utils

import (
	"os"
	"strings"
)

func CheckSnapshot() (string, error) {
	data, err := os.ReadFile(".litevc/HEAD")

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
