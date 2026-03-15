package main

import "fmt"

func main() {
	const freezeingF, boilingF = 32.0, 212.0
	fmt.Printf("freezing point = %g°F or %g°G\n", freezeingF, fToC(freezeingF))
	fmt.Printf("boiling point = %g°F or %g°G\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
