package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("0")
	fmt.Printf("%d\t", len(b))

	for {
		b = bytes.Repeat(b, 2)
		fmt.Printf("%d\t", len(b))
	}
}
