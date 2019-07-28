package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			switch os.Args[1] {
			case "complex64":
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			case "complex128":
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			case "bigfloat":
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			case "bigrat":
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex64) color.Color {
	const iterations = 200
	const contrast = 5

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255, 0, 0, 255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 5

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255, 0, 0, 255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 5

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255, 0, 0, 255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 5

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255, 0, 0, 255 - contrast*n}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
