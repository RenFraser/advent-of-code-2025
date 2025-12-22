package main

import (
	"strconv"
	"testing"
)

func TestRingBufferIteration(t *testing.T) {
	rb := RingBuffer{}
	if rb.Current != 0 {
		t.Fatal("Failed to point at zero on init")
	}

	i := 0
	for i < 100 {
		if i != rb.Current {
			t.Fatal("Failed going right. i: " + strconv.Itoa(i) + " current: " + strconv.Itoa(rb.Current))
		}
		rb.nextRight()
		i++
	}

	// current should be 100/0 now
	if rb.Current != 0 {
		t.Fatal("Failed going 99 -> 100")
	}

	if rb.FoundZeroes != 1 {
		t.Fatal("Expected 1 found zero")
	}

	rb.nextLeft()
	if rb.Current != 99 {
		t.Fatal("Failed going 0 -> 99")
	}

	i = 99
	for i >= 0 {
		if i != rb.Current {
			t.Fatal("Failed going down")
		}
		rb.nextLeft()
		i--
	}

	if rb.FoundZeroes != 2 {
		t.Fatal("Expected 2 found zeroes")
	}

}
