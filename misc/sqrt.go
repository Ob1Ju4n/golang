package golang

import (
	"fmt"
	"math"
)

// Sqrt - Calculates the square root of a given number x by using
// Newton's method
func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	for i, j := 0, 0.0; i < 10 && !equivalentFloats(j, z); i++ {
		j = z
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("Z: %.2f\t J: %.2f\n", z, j)
	}
	return z, nil
}

func equivalentFloats(a, b float64) bool {
	const threshold = 1e-9
	return math.Abs(a-b) <= threshold
}

// ErrNegativeSqrt - Custom error thrown by the Sqrt function
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot calculate Sqrt for negative number: %v", float64(e))
}
