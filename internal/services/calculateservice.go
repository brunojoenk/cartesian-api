package services

import (
	"log"
	"sort"

	"github.com/brunojoenk/cartesian-api/internal/configs"
	"github.com/brunojoenk/cartesian-api/internal/models"
)

func CalculatePoints(pointRequest models.Point) ([]models.Point, error) {
	var pointsIntoBlockWise []models.Point

	pointsOrigin, err := loadPointsOrigin(configs.GetPathFileToLoadPoints())

	if err != nil {
		log.Printf("Error on get points origin %s", err.Error())
		return nil, err
	}

	for _, pointOrigin := range pointsOrigin {
		distanceCalculated := calculateManhattanDistance(pointRequest, pointOrigin)
		if pointRequest.Distance >= distanceCalculated {
			pointsIntoBlockWise = append(pointsIntoBlockWise, *models.NewPoint(pointOrigin.X, pointOrigin.Y, distanceCalculated))
		}
	}

	sort.Slice(pointsIntoBlockWise, func(i, j int) bool {
		return pointsIntoBlockWise[i].Distance < pointsIntoBlockWise[j].Distance
	})

	return pointsIntoBlockWise, nil
}

// Following formule distance(p1,p2) = |x1-x2| + |y1-y2|
func calculateManhattanDistance(a, b models.Point) int {
	return absolute(a.X-b.X) + absolute(a.Y-b.Y)
}

// Does not have this function in lib math for int types
func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
