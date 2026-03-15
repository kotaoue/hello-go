package main

import (
	"fmt"
	"time"

	"github.com/tj/go-spin"
)

func main() {
	go spinner(100 * time.Microsecond)
	const n = 40
	fibN := fib(n) // slow
	fmt.Printf("\rFibonaccci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	s := spin.New()
	s.Set(spin.Spin3)
	for {
		fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
		time.Sleep(delay)
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
