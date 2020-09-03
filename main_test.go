package main

import "testing"

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
			Describe: "Test with steps Array 1",
			Input:    "DDUUUUDD",
			Expected: 1,
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
