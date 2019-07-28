package main

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

const (
	width, height = 1024, 1024
)

func get(q map[string][]string, k string) (string, error) {
	if len(q[k]) > 0 {
		return q[k][0], nil
	}
	return "", errors.New("the value is not found")
}

func getInt(q map[string][]string, k string, defaults int) int {
	v, err := get(q, k)
	if err != nil {
		return defaults
	}
	res, err := strconv.Atoi(v)
	if err != nil {
		return defaults
	}
	return res
}

func getFloat64(q map[string][]string, k string, defaults float64) float64 {
	v, err := get(q, k)
	if err != nil {
		return defaults
	}
	res, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return defaults
	}
	return res
}

func getString(q map[string][]string, k string, defaults string) string {
	v, err := get(q, k)
	if err != nil {
		return defaults
	}
	if v == "" {
		return defaults
	}
	return v
}

//FractalParameter is FractalParameter.
type FractalParameter struct {
	xmin float64
	ymin float64
	xmax float64
	ymax float64
}

func buildFractalParameter() FractalParameter {
	return FractalParameter{
		xmin: -2,
		ymin: -2,
		xmax: 2,
		ymax: 2,
	}
}

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.RequestURI)
		if err != nil {
			log.Fatal(err)
		}
		q := u.Query()
		p := buildFractalParameter()
		p.xmax = math.Abs(getFloat64(q, "x", 2))
		p.ymax = math.Abs(getFloat64(q, "y", 2))
		p.xmin = -1 * p.xmax
		p.ymin = -1 * p.ymax
		w.Header().Set("Content-Type", "image/png")
		fractal(w, p)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

func fractal(out io.Writer, p FractalParameter) {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(p.ymax-p.ymin) + p.ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(p.xmax-p.xmin) + p.xmin
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
