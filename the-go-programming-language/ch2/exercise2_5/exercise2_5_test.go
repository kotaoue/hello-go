// e.g. go test -bench . -benchmem
package main

import (
	"testing"

	popcount "github.com/kotaoue/the-go-programming-language/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(uint64(i))
	}
}

func BenchmarkPopCountBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountBitShift(uint64(i))
	}
}

func BenchmarkPopCountBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountBitClear(uint64(i))
	}
}
