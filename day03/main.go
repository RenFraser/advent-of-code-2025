package main

import (
	"fmt"
	"math/big"
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

// findMaxJoltageNBatteries finds the maximum n-digit joltage by selecting
// exactly n batteries from the bank. Uses a greedy monotonic stack approach.
func findMaxJoltageNBatteries(bank string, n int) string {
	if len(bank) < n || n <= 0 {
		return ""
	}

	// If we need all batteries, return the whole string
	if len(bank) == n {
		return bank
	}

	// We need to skip (len(bank) - n) batteries to keep exactly n
	toRemove := len(bank) - n
	stack := make([]byte, 0, n)

	for i := 0; i < len(bank); i++ {
		digit := bank[i]

		// While we can still remove digits and current digit is larger than stack top,
		// pop from stack to make room for the larger digit
		for len(stack) > 0 && toRemove > 0 && digit > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			toRemove--
		}

		stack = append(stack, digit)
	}

	// If we still need to remove digits (all remaining are equal/ascending),
	// remove from the end
	return string(stack[:n])
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

func runPart2(filename string) (*big.Int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	total := new(big.Int)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		joltageStr := findMaxJoltageNBatteries(line, 12)
		joltage := new(big.Int)
		joltage.SetString(joltageStr, 10)
		total.Add(total, joltage)
	}

	return total, nil
}

func main() {
	// Part 1
	result, err := run("./input/input.txt")
	if err != nil {
		fmt.Printf("Error (Part 1): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 - Total output joltage: %v\n", result)

	// Part 2
	result2, err := runPart2("./input/input.txt")
	if err != nil {
		fmt.Printf("Error (Part 2): %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 - Total output joltage: %v\n", result2)
}
