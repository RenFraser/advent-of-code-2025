package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run(filename string) (int, error) {
	content, err := os.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), ",")

	total := 0

	for _, line := range lines {
		values := strings.Split(string(line), "-")

		current, err := strconv.Atoi(values[0]) // start
		if err != nil {
			return 0, err
		}

		finish, err := strconv.Atoi(values[1])
		if err != nil {
			return 0, err
		}

		for current <= finish {
			// NOTE: for part 1, it's not recursive, just do the midpoint.
			mid := CountDigits(current) / 2
			if RepeatsAt(current, mid) {
				total += current
			}
			current++
		}
	}

	return total, nil
}

func main() {
	result, err := run("./input/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", strconv.Itoa(result))

}
