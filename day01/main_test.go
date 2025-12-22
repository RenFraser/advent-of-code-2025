package main

import (
	"strconv"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	numZeroes, err := run(50, "./input/test_input.txt")

	if err != nil {
		t.Fatal(err)
	}

	if numZeroes != 6 {
		t.Fatal("Expected 3 zeroes but found " + strconv.Itoa(numZeroes))
	}
}
