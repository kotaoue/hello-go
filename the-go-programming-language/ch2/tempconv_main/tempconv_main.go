package main

import (
	"fmt"

	tempconv "github.com/kotaoue/the-go-programming-language/ch2/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	BoilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", BoilingF-tempconv.CToF(tempconv.BoilingC))
	// fmt.Printf("%g\n", BoilingF-tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f)
	fmt.Println(c == tempconv.Celsius(f))

	c2 := tempconv.FToC(212.0)
	fmt.Println(c2.String())
	fmt.Printf("%v\n", c2)
	fmt.Printf("%s\n", c2)
	fmt.Println(c2)
	fmt.Printf("%g\n", c2)
	fmt.Println(float64(c2))
}
