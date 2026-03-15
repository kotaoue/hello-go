package main

import (
	"fmt"
	"math"
)

const e = 2.71828

// Avogadro is example for number of very large.
const Avogadro = 6.02214129e23

// Planck is example for number of very small.
const Planck = 6.62606957e-34

func main() {
	p63_1()
	p63_2()
	p63_3()
	p63_4()
	p63_5()
	p64_1()
	p64_2()
}

func p63_1() {
	fmt.Printf("%g\n", math.MaxFloat32)
	fmt.Printf("%g\n", math.MaxFloat64)
}

func p63_2() {
	var f32 float32 = 16777216
	fmt.Printf("%g\n", f32)
	fmt.Printf("%t\n", f32 == f32+1)

	var f64 float64 = 16777216
	fmt.Printf("%g\n", f64)
	fmt.Printf("%t\n", f64 == f64+1)
}

func p63_3() {
	fmt.Printf("%g\n", e)
}

func p63_4() {
	fmt.Printf("%g\n", Avogadro)
	fmt.Printf("%g\n", Planck)
	fmt.Printf("%t\n", Avogadro >= Planck)
	fmt.Printf("%t\n", Avogadro == Planck)
	fmt.Printf("%t\n", Avogadro <= Planck)
}

func p63_5() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d E^ = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func p64_1() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}

func p64_2() {
	nan := math.NaN()
	fmt.Println(nan, nan == nan, nan < nan, nan > nan)
}
