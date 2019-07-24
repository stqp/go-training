package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

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

func getString(q map[string][]string, k string, defaults string) (string, error) {
	v, err := get(q, k)
	if err != nil {
		return defaults, err
	}
	if v == "" {
		return defaults, nil
	}
	return v, nil
}

//SvgParameter is SvgParameter.
type SvgParameter struct {
	width  int
	height int
	color  string
}

func buildSvgParameter() SvgParameter {
	return SvgParameter{
		width:  600,
		height: 320,
		color:  "white",
	}
}

func main() {

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			u, err := url.Parse(r.RequestURI)
			if err != nil {
				log.Fatal(err)
			}
			q := u.Query()
			p := buildSvgParameter()
			p.height, _ = getInt(q, "height", p.height)
			p.width, _ = getInt(q, "width", p.width)
			p.color, _ = getString(q, "color", p.color)
			w.Header().Set("Content-Type", "image/svg+xml")
			svg(w, p)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

}

func svg(out io.Writer, p SvgParameter) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", p.width, p.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(out, "<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				p.color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
