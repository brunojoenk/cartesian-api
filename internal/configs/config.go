package configs

import (
	"os"
	"strings"
)

const defaultPath = "data/points.json"

func GetPathFileToLoadPoints() string {
	path := os.Getenv("PATH_DATA_POINTS")

	if strings.TrimSpace(path) == "" {
		return defaultPath
	}

	return path
}
