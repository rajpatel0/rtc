package tuple

import "math"

const Epsilon = 10e-6

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func IsPoint(tup Tuple) bool {
	return Equal(tup.W, 0.0)
}

func IsVector(tup Tuple) bool {
	return Equal(tup.W, 1.0)
}

func Equal(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}
