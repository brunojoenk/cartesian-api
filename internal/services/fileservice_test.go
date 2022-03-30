package services

import (
	"os"
	"testing"
)

var pathToLoadPointsWhenIsTest = "../../data/points.json"

func TestLoadPointsOrigin(t *testing.T) {
	os.Setenv("PATH_DATA_POINTS", pathToLoadPointsWhenIsTest)
	points, _ := loadPointsOrigin(pathToLoadPointsWhenIsTest)

	if len(points) == 0 {
		t.Error("Error on load points")
	}
}

func TestLoadPointsOriginWhenNotExistsPath(t *testing.T) {
	os.Setenv("PATH_DATA_POINTS", pathToLoadPointsWhenIsTest)
	_, err := loadPointsOrigin("anyway")

	if err == nil {
		t.Error("Must be error on get anyway file")
	}
}

func TestLoadPointsOriginWhenFileInIncompatible(t *testing.T) {
	os.Setenv("PATH_DATA_POINTS", pathToLoadPointsWhenIsTest)
	_, err := loadPointsOrigin("fileservice.go")

	if err == nil {
		t.Error("Must be error on convert fileservice.go file")
	}
}
