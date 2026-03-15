package main

import (
	"flag"
	"fmt"
)

var temp = tempconv.celsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

func CelciusFlag(name string, value Celcius, usage string) *Celcius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Scanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = f.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}
