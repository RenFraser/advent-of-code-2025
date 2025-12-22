package main

import (
	"strconv"
)

// only a sith deals in absolutes
var memoized map[int]int = map[int]int{0: 1}

func CountDigits(number int) int {
	if number < 0 {
		number = -number
	}

	var digits, ok = memoized[number]
	if ok {
		return digits
	}

	var original = number
	for number > 0 {
		number /= 10
		digits++
	}

	memoized[original] = digits

	return digits
}

// two pointers but left pointer is always at the index 0
// next is the index at which the sequence starts again
func RepeatsAt(number int, window int) bool {
	numDigits := CountDigits(number)
	if numDigits == 1 || numDigits <= window || window < 1 {
		return false
	}

	chars := []rune(strconv.Itoa(number))

	for r := window; r < len(chars); r += window {
		for l := 0; l < window; l++ {
			if l+r >= len(chars) || chars[l] != chars[r+l] {
				return false || RepeatsAt(number, window-1)
			}
		}
	}

	return true
}
