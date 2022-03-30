package routers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var router = CreateRouters()

var pathToLoadPointsWhenIsTest = "../../data/points.json"

func TestStatusOk(t *testing.T) {
	os.Setenv("PATH_DATA_POINTS", pathToLoadPointsWhenIsTest)
	request, err := http.NewRequest("GET", "/api/points?x=1&y=1&distance=25", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected code %d, received code %d", http.StatusOK, recorder.Code)
	}
}

func TestWhenQueryParamXIsInvalid(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/points?x=a&y=1&distance=25", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected code %d, received code %d", http.StatusBadRequest, recorder.Code)
	}

	var messageReturned string
	json.Unmarshal(recorder.Body.Bytes(), &messageReturned)
	messageExpected := "Invalid query parameter: x"
	if strings.Compare(messageReturned, messageExpected) != 0 {
		t.Errorf("Expected %s, received %s", messageExpected, messageReturned)
	}
}

func TestWhenQueryParamYIsInvalid(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/points?x=1&y=a&distance=25", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected code %d, received code %d", http.StatusBadRequest, recorder.Code)
	}

	var messageReturned string
	json.Unmarshal(recorder.Body.Bytes(), &messageReturned)
	messageExpected := "Invalid query parameter: y"
	if strings.Compare(messageReturned, messageExpected) != 0 {
		t.Errorf("Expected %s, received %s", messageExpected, messageReturned)
	}
}

func TestWhenQueryParamDistanceIsInvalid(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/points?x=1&y=1&distance=a", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected %d, received %d", http.StatusBadRequest, recorder.Code)
	}

	var messageReturned string
	json.Unmarshal(recorder.Body.Bytes(), &messageReturned)
	messageExpected := "Invalid query parameter: distance"
	if strings.Compare(messageReturned, messageExpected) != 0 {
		t.Errorf("Expected %s, received %s", messageExpected, messageReturned)
	}
}
