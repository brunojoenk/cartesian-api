package services

import (
	"os"
	"testing"

	"github.com/brunojoenk/cartesian-api/internal/models"
)

func TestManhattanDistanceWithValuesFromPointsJson(t *testing.T) {
	os.Setenv("PATH_DATA_POINTS", pathToLoadPointsWhenIsTest)
	pointRequested := models.Point{X: 1, Y: 1, Distance: 20}
	expectedPoints := []models.Point{
		{X: 2, Y: -8, Distance: 10},
		{X: -6, Y: -10, Distance: 18},
	}

	pd, _ := CalculatePoints(pointRequested)

	for i := 0; i < len(pd); i++ {
		if pd[i].X != expectedPoints[i].X {
			t.Errorf("Error on comparate points X Calculated:%v  X Expected: %v", pd[i].X, expectedPoints[i].X)
		}
		if pd[i].Y != expectedPoints[i].Y {
			t.Errorf("Error on comparate points Y Calculated:%v  Y Expected: %v", pd[i].Y, expectedPoints[i].Y)
		}
		if pd[i].Distance != expectedPoints[i].Distance {
			t.Errorf("Error on comparate points Distance Calculated:%v Distance Expected: %v", pd[i].Distance, expectedPoints[i].Distance)
		}
	}
}
