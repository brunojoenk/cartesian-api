package services

import (
	"encoding/json"
	"log"
	"os"

	"github.com/brunojoenk/cartesian-api/internal/models"
)

func loadPointsOrigin(path string) ([]models.Point, error) {
	byteValue, err := os.ReadFile(path)

	if err != nil {
		log.Printf("Can not read the file: %s, error: %s", path, err.Error())
		return nil, err
	}

	var points []models.Point
	err = json.Unmarshal(byteValue, &points)

	if err != nil {
		log.Println("Can not convert data of file to Point struct", err)
		return nil, err
	}

	return points, nil
}
