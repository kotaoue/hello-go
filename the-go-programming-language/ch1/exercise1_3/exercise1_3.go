// e.g. go run exercise1_3.go a=1 b=2 c=3 d=4 e=5
// a=1 b=2 c=3 d=4 e=5
// 0.000036s
// a=1 b=2 c=3 d=4 e=5
// 0.000058s
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("%3fs\n", secs)

	fmt.Println(strings.Join(os.Args[1:], " "))
	secs = time.Since(start).Seconds()
	fmt.Printf("%3fs\n", secs)
}
