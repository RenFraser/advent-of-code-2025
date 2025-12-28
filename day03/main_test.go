package main

import "testing"

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
