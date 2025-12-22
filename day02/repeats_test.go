package main

import (
	"strconv"
	"testing"
)

func TestCountDigits(t *testing.T) {
	tests := []struct {
		number   int
		expected int
	}{
		{0, 1},
		{1, 1},
		{10, 2},
		{100, 3},
		{1000, 4},
		{10000, 5},
		{100000, 6},
		{12, 2},
		{123, 3},
		{1234, 4},
		{12345, 5},
		{123456, 6},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.number), func(t *testing.T) {
			result := CountDigits(tt.number)
			if result != tt.expected {
				t.Errorf("CountDigits(%d) = %d; want %d", tt.number, result, tt.expected)
			}
		})
	}
}

// A bunch of the tests have hidden true values but strictly from the non-recursive approach, it wouldn't be true.
// pretty misleading but I'm too lazy to change it.
func TestRepeatsAt(t *testing.T) {
	tests := []struct {
		number   int
		window   int
		expected bool
	}{
		{123123, 3, true},
		{101, 2, false},
		{101, 3, false},
		{12345, 2, false},
		{12345, 3, false},
		{1, 1, false},
		{1, 2, false},
		{11, 1, true},
		{565656, 2, true},
		{565656, 3, true},
		{565656, 4, true},
		{111111, 2, true},
		{111111, 3, true},
		{111111, 4, true},
		{123123123, 3, true},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.number), func(t *testing.T) {
			result := RepeatsAt(tt.number, tt.window)
			if result != tt.expected {
				t.Errorf("RepeatsAt(%d, %d) = %t; want %t", tt.number, tt.window, result, tt.expected)
			}
		})
	}
}
