package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	result, err := run("./input/test_input.txt")
	if err != nil {
		fmt.Println(err)
		t.Fatalf("Got %d, expected: 1227775554", result)
	}
}
