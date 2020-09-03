package main

import (
	"reflect"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	type Test struct {
		Describe string
		Input    []int
		Expected int
	}

	tests := []Test{
		{
			Describe: "Test with Array 1",
			Input:    []int{1, 2, 1, 2, 3, 2},
			Expected: 2,
		},
		{
			Describe: "Test with Array 2",
			Input:    []int{10, 20, 20, 10, 10, 30, 50, 10, 20},
			Expected: 3,
		},
		{
			Describe: "Test with no pairs Array",
			Input:    []int{1, 2, 3, 4, 5, 6},
			Expected: 0,
		},
	}

	for _, test := range tests {
		result := sockMerchant(test.Input)
		if result == test.Expected {
			t.Logf("\n%s => SUCCESS, expected: %d; result: %d", test.Describe, test.Expected, result)
		} else {
			t.Errorf("\n%s => FAILED, expected: %d; result: %d", test.Describe, test.Expected, result)
		}
	}
}

func TestCountingValleys(t *testing.T) {
	type Test struct {
		Describe string
		Input    string
		Expected int
	}

	tests := []Test{
		{
			Describe: "Test steps with Array 1",
			Input:    "DDUUUUDD",
			Expected: 1,
		},
		{
			Describe: "Test steps with Array 2",
			Input:    "UDDDUDUU",
			Expected: 1,
		},
		{
			Describe: "Test steps through 3 valleys",
			Input:    "DUDUDU",
			Expected: 3,
		},
		{
			Describe: "Test steps without valleys",
			Input:    "UDUDUD",
			Expected: 0,
		},
	}

	for _, test := range tests {
		result := countingValleys(test.Input)
		if result == test.Expected {
			t.Logf("\n%s => SUCCESS, expected: %d; result: %d", test.Describe, test.Expected, result)
		} else {
			t.Errorf("\n%s => FAILED, expected: %d; result: %d", test.Describe, test.Expected, result)
		}
	}
}

func TestNominalLayers(t *testing.T) {
	type Test struct {
		Describe string
		Input    int
		Expected []int
	}

	tests := []Test{
		{
			Describe: "Test millions nominal",
			Input:    1345679,
			Expected: []int{1000000, 300000, 40000, 5000, 600, 70, 9},
		},
		{
			Describe: "Test thousands nominal",
			Input:    4758,
			Expected: []int{4000, 700, 50, 8},
		},
		{
			Describe: "Test hundreds nominal",
			Input:    214,
			Expected: []int{200, 10, 4},
		},
		{
			Describe: "Test dozens nominal",
			Input:    27,
			Expected: []int{20, 7},
		},
		{
			Describe: "Test units nominal",
			Input:    7,
			Expected: []int{7},
		},
	}

	for _, test := range tests {
		result := nominalLayers(test.Input)
		if reflect.DeepEqual(result, test.Expected) {
			t.Logf("\n%s => SUCCESS, expected: %v; result: %v", test.Describe, test.Expected, result)
		} else {
			t.Errorf("\n%s => FAILED, expected: %v; result: %v", test.Describe, test.Expected, result)
		}
	}
}
