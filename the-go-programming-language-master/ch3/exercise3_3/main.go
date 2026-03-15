// $ go run ch3/exercise3_3/main.go > ch3/exercise3_3/surface.svg
// $ open -a '/Applications/Google Chrome.app' ch3/exercise3_3/surface.svg
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width; 0.7' "+
		"width='%d' height='%d'>", width, height)
	// zmax := 0.0
	// zmin := 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			/*
				z := (az + bz + cz + dz) / 4

				if zmax < z {
					zmax = z
				}
				if zmin > z {
					zmin = z
				}
			*/

			if !isNaN([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, fill(az, bz, cz, dz))
			}
		}
	}
	// fmt.Println(zmax) = 0.9557691306147335
	// fmt.Println(zmin) =
	fmt.Println("</svg>")

}

func isNaN(f []float64) bool {
	for _, v := range f {
		if math.IsNaN(v) {
			return true
		}
	}
	return false
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func fill(az, bz, cz, dz float64) (c string) {
	z := (az + bz + cz + dz) / 4

	if z >= 0.9 {
		// 0.9557691306147335
		c = "#ff0000"
	} else if z <= -0.2 {
		// -0.2147958119649517
		c = "#0000ff"
	}
	return
}
