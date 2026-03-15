package main

import (
	"fmt"

	popcount "github.com/kotaoue/the-go-programming-language/ch2/popcount"
)

func main() {
	for i := 0; i < 64; i++ {
		r1 := popcount.PopCount(uint64(i))
		r2 := popcount.PopCountLoop(uint64(i))
		r3 := popcount.PopCountBitShift(uint64(i))
		r4 := popcount.PopCountBitClear(uint64(i))
		fmt.Printf("%d %d %d %d %d\n", i, r1, r2, r3, r4)
	}
}
