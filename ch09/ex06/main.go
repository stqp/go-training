package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"sync"
)

func main() {}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 10240, 1024
)

func runConcurrencyMandelbrot(out io.Writer, cuncurrency int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup
	c := make(chan struct{}, cuncurrency)

	for py := 0; py < height; py++ {
		wg.Add(1)
		go func(py int) {
			c <- struct{}{}
			defer func() {
				<-c
				wg.Done()
			}()
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)

				img.Set(px, py, mandelbrot(z))
			}
		}(py)
	}
	wg.Wait()
	png.Encode(out, img)
}

func runHeavyMandelbrot(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
