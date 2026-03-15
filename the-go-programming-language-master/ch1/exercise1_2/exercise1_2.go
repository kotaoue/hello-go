// e.g go run exercise1_2.go a b c d e f g h i j k
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}
