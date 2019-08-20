package liftgotofloor_test

import "testing"

func Test_GoToFloor_Input_GoToFloor_G_Should_Be_GoToFloor_G(t *testing.T) {
	expectedResult := "G"
	floor := "G"

	actualResult := GoToFloor(floor)

	if expectedResult != actualResult {
		t.Errorf("Expect %s but got %s", expectedResult, actualResult)
	}
}
