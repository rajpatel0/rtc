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
	return Equal(tup.W, 1.0)
}

func IsVector(tup Tuple) bool {
	return Equal(tup.W, 0.0)
}

func Equal(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func EqualTuple(tup1, tup2 Tuple) bool {
	return Equal(tup1.X, tup2.X) && Equal(tup1.Y, tup2.Y) && Equal(tup1.Z, tup2.Z) && Equal(tup1.W, tup2.W)
}
