package models

type Point struct {
	X        int `json:"x"`
	Y        int `json:"y"`
	Distance int `json:"distance"`
}

func NewPoint(x, y, distance int) Point {
	return Point{
		X:        x,
		Y:        y,
		Distance: distance,
	}
}
