package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func cSqrt(x float64) (float64, error) {
	if (x == 0) {
		return 0, nil
	} else if x < 0 {
		return math.NaN(), ErrNegativeSqrt(x)
	}
	
	z := 1.0
	
	for deviation := ((z*z - x) / (2*z)); (math.Abs(deviation) >= 0.001); deviation = ((z*z - x) / (2*z)) {
		z -= deviation
	}
	return z, nil
}

func main() {
	fmt.Println(cSqrt(2))
	if _, err := cSqrt(-2); err != nil {
		fmt.Println(err)
	}
}
