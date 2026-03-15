package main

import "fmt"

func main() {
	sum := 1
	for sum < 65535 {
		sum += sum
	}
	fmt.Println(sum)
}
