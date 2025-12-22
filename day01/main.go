package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run(start int, filename string) (int, error) {
	rb := &RingBuffer{start, 0}
	content, err := os.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	var lines = strings.Split(string(content), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var direction = rune(line[0])
		var distance, err = strconv.Atoi(line[1:])

		if err != nil {
			return 0, err
		}

		switch direction {
		case 'L':
			i := 0
			for i < distance {
				rb.nextLeft()
				i++
			}
		case 'R':
			i := 0
			for i < distance {
				rb.nextRight()
				i++
			}
		default:
			return 0, errors.New("Unexpected direction: " + string(direction))
		}
	}

	return rb.FoundZeroes, nil
}

func main() {
	var result, err = run(50, "./input/input.txt")
	if err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Number of zeroes: %v\n", result)
	}

}
