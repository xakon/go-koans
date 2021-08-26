package go_koans

func divideFourBy(i int) (n int) {
	defer func() {
		if recover() != nil {
			n = 2
		}
	}()

	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	assert(true) // panics are exceptional errors at runtime

	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
