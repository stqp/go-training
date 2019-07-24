package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 300            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type zfunc func(x, y float64) float64

func egg(x, y float64) float64 {
	return math.Sin(math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2)))
}
func mogul(x, y float64) float64 {
	return math.Sin(math.Sin(x)*math.Cos(y)) / 3
	//return (3*x + math.Pow(y, 2)) * math.Abs(math.Sin(x)+math.Cos(y))
}
func saddle(x, y float64) float64 {
	return math.Pow(x, 2) - math.Pow(y, 2)
}

func main() {
	var f zfunc
	if os.Args[1] == "egg" {
		f = egg
	} else if os.Args[1] == "mogul" {
		f = mogul
	} else if os.Args[1] == "saddle" {
		f = saddle
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func isFinite(xs []float64) bool {
	for _, x := range xs {
		if math.Abs(x) == math.Inf(0) {
			return false
		}
	}
	return true
}

func corner(i, j int, f zfunc) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
