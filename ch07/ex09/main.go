package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type TracksSortedResult struct {
	Items []*Track
}

const (
	title = iota
	artist
	album
	year
	length
)

var fields = map[string]int{
	"title":  0,
	"artist": 1,
	"album":  2,
	"year":   3,
	"length": 4,
}

type MultiSortWidget struct {
	t []*Track
	s []int
	f [5]func(x, y *Track) (bool, bool)
}

func (x MultiSortWidget) Len() int           { return len(x.t) }
func (x MultiSortWidget) Less(i, j int) bool { return x.less(i, j) }
func (x MultiSortWidget) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x MultiSortWidget) less(i, j int) bool {
	for k := len(x.s) - 1; k >= 0; k-- {
		v := x.s[k]
		res, decided := x.f[v](x.t[i], x.t[j])
		if decided {
			return res
		}
	}
	return false
}

func addSortTargetFiled(fields []int, field int) []int {
	for i, f := range fields {
		if f == field {
			fields = append(fields[:i], fields[i+1:]...)
			fields = append(fields, field)
			return fields
		}
	}
	fields = append(fields, field)
	return fields
}

func main() {
	widget := MultiSortWidget{
		t: tracks,
	}
	widget.f[title] = func(x, y *Track) (res bool, decided bool) {
		if x.Title != y.Title {
			return x.Title < y.Title, true
		}
		return false, false
	}
	widget.f[artist] = func(x, y *Track) (res bool, decided bool) {
		if x.Artist != y.Artist {
			return x.Artist < y.Artist, true
		}
		return false, false
	}
	widget.f[album] = func(x, y *Track) (res bool, decided bool) {
		if x.Album != y.Album {
			return x.Album < y.Album, true
		}
		return false, false
	}
	widget.f[year] = func(x, y *Track) (res bool, decided bool) {
		if x.Year != y.Year {
			return x.Year < y.Year, true
		}
		return false, false
	}
	widget.f[length] = func(x, y *Track) (res bool, decided bool) {
		if x.Length != y.Length {
			return x.Length < y.Length, true
		}
		return false, false
	}

	http.HandleFunc("/", middle(&widget))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func middle(widget *MultiSortWidget) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.URL.String())
		query := u.Query()
		for k, v := range fields {
			if query.Get(k) != "" {
				widget.s = addSortTargetFiled(widget.s, v)
				break
			}
		}
		//ユーザ入力に沿って並び変える.
		sort.Sort(widget)
		data := TracksSortedResult{Items: widget.t}
		if err := mytemplate().Execute(w, data); err != nil {
			log.Fatal(err)
		}
	}
}

func mytemplate() *template.Template {
	body := `
	<html>
	<body>
		<form>
			<table>
				<tr>
					<th><input type="submit" name="title"  value="Title"></th>
					<th><input type="submit" name="artist" value="Artist"></th>
					<th><input type="submit" name="album"  value="Album"></th>
					<th><input type="submit" name="year"   value="Year"></th>
					<th><input type="submit" name="length" value="Length"></th>
				</tr>
				{{range .Items}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
				</tr>
				{{end}}
			</table>
		</form>
	</body>
	</html>
	`
	return template.Must(template.New("tracktable").Parse(body))
}
