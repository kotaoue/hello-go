package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2.0 * z)

		fmt.Printf("%d : %g\n", i, z)
		if math.Abs(math.Sqrt(x)-z) < 0.00000000001 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
