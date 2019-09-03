package main

import (
	"testing"
)

var (
	//	url = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"
	url = "https://web.ics.purdue.edu/~gchopra/class/public/pages/webdesign/05_simple.html"
)

func TestMain(t *testing.T) {

	expWords := 197
	expImages := 2

	words, images, err := CountWordsAndImages(url)
	if err != nil {
		t.Errorf("Error catch. %v\n", err)
	}
	if words != expWords {
		t.Error("Words..  Expected:", expWords, "Actual:", words)
	}
	if images != expImages {
		t.Error("images..  Expected:", expImages, "Actual:", images)
	}

}
