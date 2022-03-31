package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brunojoenk/cartesian-api/internal/models"
	"github.com/brunojoenk/cartesian-api/internal/services"
	"github.com/pkg/errors"
)

//The tests of handlers are embedded on routers_test.go class

func HandleDistance(writer http.ResponseWriter, request *http.Request) {
	pointRequest, err := getParameters(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err.Error())
		return
	}

	points, err := services.CalculatePoints(*pointRequest)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(points)
}

func getParameters(request *http.Request) (*models.Point, error) {
	queryPath := request.URL.Query()
	x, err := strconv.Atoi(queryPath.Get("x"))
	if err != nil {
		return nil, errors.New("Invalid query parameter: x")
	}

	y, err := strconv.Atoi(queryPath.Get("y"))
	if err != nil {
		return nil, errors.New("Invalid query parameter: y")
	}

	distance, err := strconv.Atoi(queryPath.Get("distance"))
	if err != nil {
		return nil, errors.New("Invalid query parameter: distance")
	}

	return models.NewPoint(x, y, distance), nil

}

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("Welcome to cartesian API")
}
