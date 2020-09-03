package main

import "testing"

func TestAlgorithmOne(t *testing.T) {
	socks := 7
	colors := []int{1, 2, 1, 2, 3, 2}
	expected := 2

	result := algorithmNumberOne(socks, colors)

	if result == expected {
		t.Logf("Algorithm Number One SUCCESS, expected %d and the result is %d", result, expected)
	}

	t.Errorf("Algorithm Number One FAILED, expected %d but the result is %d", result, expected)
}
