package main

import (
	"fmt"
	"math"
)

func cSqrt(x float64) float64 {
	if (x == 0) {
		return 0
	} else if x < 0 {
		return math.NaN()
	}
	
	z := 1.0
	
	for deviation := ((z*z - x) / (2*z)); (math.Abs(deviation) >= 0.001); deviation = ((z*z - x) / (2*z)) {
		z -= deviation
	}
	return z
}

func main() {

	const value = 2
	fmt.Println(cSqrt(value), math.Sqrt(value))
}
