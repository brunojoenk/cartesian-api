package utils

import (
	"testing"
)

func TestConcatStrings(t *testing.T) {
	expectedString := "Cartesian API"
	returnedString := ConcatStrings("Cartesian", " ", "API")

	if expectedString != returnedString {
		t.Errorf("Error on concat strings. Expected: %s , Returned: %s", expectedString, returnedString)
	}
}
