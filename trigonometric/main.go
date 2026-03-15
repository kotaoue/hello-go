package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	size := 360
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size}})

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			img.Set(x, y, color.White)
		}
	}

	for x := 0; x < size; x++ {
		r1 := float64(x) * math.Pi / float64(180)
		r2 := float64(x+120) * math.Pi / float64(180)
		r3 := float64(x+240) * math.Pi / float64(180)

		y1 := int(math.Round(float64(size/2)*math.Sin(r1))) + size/2
		y2 := int(math.Round(float64(size/2)*math.Sin(r2))) + size/2
		y3 := int(math.Round(float64(size/2)*math.Sin(r3))) + size/2

		fmt.Printf("%3d:%4d:%4d:%4d\n", x, y1, y2, y3)

		img.Set(x, size/2, color.RGBA{128, 128, 128, 255})
		img.Set(x, y1, color.RGBA{255, 0, 0, 255})
		img.Set(x, y2, color.RGBA{0, 255, 0, 255})
		img.Set(x, y3, color.RGBA{0, 0, 255, 255})
	}

	f, err := os.Create("image.png")
	if err != nil {
		return err
	}

	return png.Encode(f, img)
}
