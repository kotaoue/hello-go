package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z float32
}

func main() {
	fmt.Println(Vertex{1, 2, 0.25})
}
