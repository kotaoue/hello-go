package main

import (
	"testing"
)

func BenchmarkMakeRuneMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeRuneMap(100)
	}
}

func BenchmarkMake2RuneMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Make2RuneMap(100)
	}
}

func BenchmarkMakeStringMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeStringMap(100)
	}
}
