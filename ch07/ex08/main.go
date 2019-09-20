package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

const (
	title = iota
	artist
	album
	year
	length
)

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

func addClick(fields []int, field int) []int {
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

	// ユーザがクリックしたフィールドを追加する。
	widget.s = addClick(widget.s, year)
	widget.s = addClick(widget.s, title)

	banner("Before:")
	printTracks(tracks)

	sort.Sort(widget)

	banner("After:")
	printTracks(tracks)
}

func banner(title string) {
	fmt.Println()
	sep := "#"
	for i := 0; i < 20; i++ {
		fmt.Print(sep)
	}
	fmt.Println()
	fmt.Println(sep + "   " + title)
	for i := 0; i < 20; i++ {
		fmt.Print(sep)
	}
	fmt.Println()
	fmt.Println()
}
