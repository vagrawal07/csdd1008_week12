package main

import "testing"

func TestFactorial(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, tc := range testCases {
		result := Factorial(tc.n)
		if result != tc.expected {
			t.Errorf("Factorial(%d) = %d; expected %d", tc.n, result, tc.expected)
		}
	}
}
