package main

// It's not really a ring buffer. Just the idea.
// I can't be bothered to change the name now
type RingBuffer struct {
	Current     int
	FoundZeroes int
}

func (rb *RingBuffer) nextLeft() {
	if rb.Current == 0 {
		rb.Current = 99
	} else {
		rb.Current--
	}

	if rb.Current == 0 {
		rb.FoundZeroes++
	}
}

func (rb *RingBuffer) nextRight() {
	if rb.Current == 99 {
		rb.Current = 0
	} else {
		rb.Current++
	}

	if rb.Current == 0 {
		rb.FoundZeroes++
	}
}
