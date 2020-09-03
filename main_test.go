package main

import "testing"

func TestAlgorithmOne(t *testing.T) {
	colors := []int{1, 2, 1, 2, 3, 2}
	expected := 2

	result := algorithmNumberOne(colors)

	if result == expected {
		t.Logf("Algorithm Number One SUCCESS, expected %d and the result is %d", result, expected)
	} else {
		t.Errorf("Algorithm Number One FAILED, expected %d but the result is %d", result, expected)
	}

}
