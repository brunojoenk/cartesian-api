package configs

import (
	"os"
	"strings"
)

const pathDefault = "data/points.json"

func GetPathFileToLoadPoints() string {
	path := os.Getenv("PATH_DATA_POINTS")

	if strings.TrimSpace(path) == "" {
		return pathDefault
	}

	return path
}
