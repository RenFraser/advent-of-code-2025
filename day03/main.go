package main

import (
	"fmt"
	"os"
	"strings"
)

// Note: len(bank) gives us byte count, not character count. This works fine here
// since we're only dealing with ASCII digits (1-9), which are 1 byte each in UTF-8.
func findMaxJoltage(bank string) int {
	if len(bank) < 2 {
		return 0
	}

	// First pass: find the highest digit that's not in the last position
	maxTensDigit := rune('0')
	maxTensPos := -1

	for i := 0; i < len(bank)-1; i++ {
		if rune(bank[i]) > maxTensDigit {
			maxTensDigit = rune(bank[i])
			maxTensPos = i
		}
	}

	// Second pass: find the highest digit after maxTensPos
	maxOnesDigit := rune('0')

	for i := maxTensPos + 1; i < len(bank); i++ {
		if rune(bank[i]) > maxOnesDigit {
			maxOnesDigit = rune(bank[i])
		}
	}

	// Convert to two-digit number
	tens := int(maxTensDigit - '0')
	ones := int(maxOnesDigit - '0')

	return tens*10 + ones
}

func run(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(content), "\n")
	total := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		joltage := findMaxJoltage(line)
		total += joltage
	}

	return total, nil
}

func main() {
	result, err := run("./input/input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total output joltage: %v\n", result)
}
