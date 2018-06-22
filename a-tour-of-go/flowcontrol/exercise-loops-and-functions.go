package main

import (
	"fmt"
	"math"
)

const Delta = 0.0001

func isConverged(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	if d < Delta {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		tmp = z - (z*z-x)/2*z
		if d := tmp - z; isConverged(d) {
			return tmp
		}
		z = tmp
	}
	return z
}

func main() {
	attempt := Sqrt(2)
	expected := math.Sqrt(2)
	fmt.Printf("attempt = %g (expected = %g) error = %g\n",
		attempt, expected, attempt-expected)
}
