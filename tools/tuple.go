package tools

import (
	"math"
)

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

func AddTuple(tup1, tup2 Tuple) Tuple {
	return Tuple{tup1.X + tup2.X, tup1.Y + tup2.Y, tup1.Z + tup2.Z, tup1.W + tup2.W}
}

func SubTuple(tup1, tup2 Tuple) Tuple {
	return Tuple{tup1.X - tup2.X, tup1.Y - tup2.Y, tup1.Z - tup2.Z, tup1.W - tup2.W}
}

func NegateTuple(tup1 Tuple) Tuple {
	return Tuple{-tup1.X, -tup1.Y, -tup1.Z, -tup1.W}
}

func MultiplyScaler(tup1 Tuple, a float64) Tuple {
	return Tuple{tup1.X * a, tup1.Y * a, tup1.Z * a, tup1.W * a}
}

func DivideScaler(tup1 Tuple, a float64) Tuple {
	return Tuple{tup1.X / a, tup1.Y / a, tup1.Z / a, tup1.W / a}
}

func Magnitude(tup1 Tuple) float64 {
	return math.Sqrt(tup1.X*tup1.X + tup1.Y*tup1.Y + tup1.Z*tup1.Z + tup1.W*tup1.W)
}

func Normalize(tup1 Tuple) Tuple {
	mag := Magnitude(tup1)
	return DivideScaler(tup1, mag)
}

func Dot(tup1 Tuple, tup2 Tuple) float64 {
	return tup1.X*tup2.X + tup1.Y*tup2.Y + tup1.Z*tup2.Z + tup1.W*tup2.W
}

func Cross(tup1 Tuple, tup2 Tuple) Tuple {
	x := tup1.Y*tup2.Z - tup1.Z*tup2.Y
	y := tup1.Z*tup2.X - tup1.X*tup2.Z
	z := tup1.X*tup2.Y - tup1.Y*tup2.X
	return Vector(x, y, z)
}

func Color(red, green, blue float64) Tuple {
	return Vector(red, green, blue)
}

func Hadamard(color1, color2 Tuple) Tuple {
	r := color1.X * color2.X
	b := color1.Y * color2.Y
	g := color1.Z * color2.Z
	return Vector(r, b, g)
}
