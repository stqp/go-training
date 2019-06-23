package main

import (
	"errors"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func get(q map[string][]string, k string) (string, error) {
	if len(q[k]) > 0 {
		return q[k][0], nil
	}
	return "", errors.New("the value is not found")
}

func getInt(q map[string][]string, k string, defaults int) (int, error) {
	v, err := get(q, k)
	if err != nil {
		return defaults, err
	}
	res, err := strconv.Atoi(v)
	if err != nil {
		return defaults, err
	}
	return res, nil
}

func getFloat64(q map[string][]string, k string, defaults float64) (float64, error) {
	v, err := get(q, k)
	if err != nil {
		return defaults, err
	}
	res, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return defaults, err
	}
	return res, nil
}

func getString(q map[string][]string, k string) (string, error) {
	return get(q, k)
}

// LissajousParameter is parameter
type LissajousParameter struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func buildLissajousParameter() LissajousParameter {
	return LissajousParameter{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	p := buildLissajousParameter()

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			u, err := url.Parse(r.RequestURI)
			if err != nil {
				log.Fatal(err)
			}
			q := u.Query()
			p.cycles, _ = getInt(q, "cycles", p.cycles)
			p.res, _ = getFloat64(q, "res", p.res)
			p.size, _ = getInt(q, "size", p.size)
			p.nframes, _ = getInt(q, "nframes", p.nframes)
			p.delay, _ = getInt(q, "delay", p.delay)
			lissajous(w, p)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, p)
}

func lissajous(out io.Writer, p LissajousParameter) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.cycles)*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
