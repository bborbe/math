package math

import "math"

func Sqrt[T ~float64](x T) T {
	return T(math.Sqrt(float64(x)))
}
