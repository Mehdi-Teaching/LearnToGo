package main

import (
	"fmt"
	"testing"
)

// IMPORTANT:
/*
	1. File name must end with [_test]. in this case min_test.go
	2. Test functions must start with [Test]. for example TestMinValue
	3. Test functions must have [t *testing.T] as parameter.
	4. run tests using command [go test [PACKAGE]]
*/

// use a struct to simplify testing
type testUnit struct {
	x, y, expected int
}

// function that is going to be tested
func minValue(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// test function
func TestMinValue(t *testing.T) {
	// table style tests
	tests := []testUnit{
		{2, 1, 1},
		{1, 0, 0},
		{2, -2, -2},
		{-1, 0, -1},
	}

	for _, unit := range tests {
		// each test need a name
		testname := fmt.Sprintf("%d,%d", unit.x, unit.y)
		// running tests using t.Run(string, func(t *testing.T) ) signature.
		t.Run(testname, testMinValue(unit))
	}
}

// if error happen in the return func, test fail otherwise test pass
func testMinValue(unit testUnit) func(t *testing.T) {
	return func(t *testing.T) {
		ans := minValue(unit.x, unit.y)

		if ans != unit.expected {
			t.Errorf("minValue(%d,%d) = %d; expect %d", unit.x, unit.y, ans, unit.expected)
		}
	}
}
