package main

import (
	"fmt"

	"github.com/kotaoue/go-interface/packages/calc"
)

func main() {
	c := &calc.Calculator{Processor: &calc.CALU{}}
	fmt.Println(c.Calc(1, 2))

	c = calc.NewCalculator(&calc.CALU{})
	fmt.Println(c.Calc(1, 2))
}
