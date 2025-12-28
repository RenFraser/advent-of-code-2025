package main

import (
	"math/big"
	"os"
	"testing"
)

func TestFindMaxJoltage(t *testing.T) {
	tests := []struct {
		name     string
		bank     string
		expected int
	}{
		{
			name:     "Example bank 1",
			bank:     "987654321111111",
			expected: 98,
		},
		{
			name:     "Example bank 2",
			bank:     "811111111111119",
			expected: 89,
		},
		{
			name:     "Example bank 3",
			bank:     "234234234234278",
			expected: 78,
		},
		{
			name:     "Example bank 4",
			bank:     "818181911112111",
			expected: 92,
		},
		{
			name:     "Simple ascending",
			bank:     "123",
			expected: 23,
		},
		{
			name:     "Simple descending",
			bank:     "321",
			expected: 32,
		},
		{
			name:     "All same digits",
			bank:     "5555",
			expected: 55,
		},
		{
			name:     "Two digits only",
			bank:     "42",
			expected: 42,
		},
		{
			name:     "Max at end",
			bank:     "12349",
			expected: 49,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxJoltage(tt.bank)
			if result != tt.expected {
				t.Errorf("findMaxJoltage(%q) = %d, expected %d", tt.bank, result, tt.expected)
			}
		})
	}
}

func TestFindMaxJoltageEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		bank     string
		expected int
	}{
		{
			name:     "Empty string",
			bank:     "",
			expected: 0,
		},
		{
			name:     "Single digit",
			bank:     "5",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxJoltage(tt.bank)
			if result != tt.expected {
				t.Errorf("findMaxJoltage(%q) = %d, expected %d", tt.bank, result, tt.expected)
			}
		})
	}
}

func TestFindMaxJoltageNBatteries(t *testing.T) {
	tests := []struct {
		name     string
		bank     string
		n        int
		expected string
	}{
		{
			name:     "Part 2 Example 1",
			bank:     "987654321111111",
			n:        12,
			expected: "987654321111",
		},
		{
			name:     "Part 2 Example 2",
			bank:     "811111111111119",
			n:        12,
			expected: "811111111119",
		},
		{
			name:     "Part 2 Example 3",
			bank:     "234234234234278",
			n:        12,
			expected: "434234234278",
		},
		{
			name:     "Part 2 Example 4",
			bank:     "818181911112111",
			n:        12,
			expected: "888911112111",
		},
		{
			name:     "Select all",
			bank:     "123",
			n:        3,
			expected: "123",
		},
		{
			name:     "Select 2 from part 1 examples",
			bank:     "987654321111111",
			n:        2,
			expected: "98",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxJoltageNBatteries(tt.bank, tt.n)
			if result != tt.expected {
				t.Errorf("findMaxJoltageNBatteries(%q, %d) = %s, expected %s", tt.bank, tt.n, result, tt.expected)
			}
		})
	}
}

func TestRunPart2(t *testing.T) {
	// Test with the example input
	example := `987654321111111
811111111111119
234234234234278
818181911112111`

	// Create temp file for testing
	tmpfile := t.TempDir() + "/test_input.txt"
	if err := os.WriteFile(tmpfile, []byte(example), 0644); err != nil {
		t.Fatal(err)
	}

	result, err := runPart2(tmpfile)
	if err != nil {
		t.Fatalf("runPart2 failed: %v", err)
	}

	expected := new(big.Int)
	expected.SetString("3121910778619", 10)

	if result.Cmp(expected) != 0 {
		t.Errorf("runPart2() = %s, expected %s", result.String(), expected.String())
	}
}
