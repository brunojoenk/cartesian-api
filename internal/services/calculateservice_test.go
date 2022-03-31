package services

import (
	"os"
	"testing"

	"github.com/brunojoenk/cartesian-api/internal/models"
)

func TestManhattanDistanceWithValuesFromPointsJson(t *testing.T) {
	os.Setenv("PATH_DATA_POINTS", pathToLoadPointsWhenIsTest)
	pointToCalculate := models.Point{X: 1, Y: 1, Distance: 20}
	expectedPoints := []models.Point{
		{X: 2, Y: -8, Distance: 10},
		{X: -6, Y: -10, Distance: 18},
	}

	points, _ := CalculatePoints(pointToCalculate)

	for i := 0; i < len(points); i++ {
		if points[i].X != expectedPoints[i].X {
			t.Errorf("Error on comparate points X Calculated:%v  X Expected: %v", points[i].X, expectedPoints[i].X)
		}
		if points[i].Y != expectedPoints[i].Y {
			t.Errorf("Error on comparate points Y Calculated:%v  Y Expected: %v", points[i].Y, expectedPoints[i].Y)
		}
		if points[i].Distance != expectedPoints[i].Distance {
			t.Errorf("Error on comparate points Distance Calculated:%v Distance Expected: %v", points[i].Distance, expectedPoints[i].Distance)
		}
	}
}

func TestManhattanDistanceWhenTryIncorrectFile(t *testing.T) {
	os.Setenv("PATH_DATA_POINTS", "anyway")
	pointRequested := models.Point{X: 1, Y: 1, Distance: 20}

	_, err := CalculatePoints(pointRequested)

	if err == nil {
		t.Errorf("Expected error")
	}
}
