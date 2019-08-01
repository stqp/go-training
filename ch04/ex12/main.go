package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	minGetAndStoreBookNumberRange = 1
	maxGetAndStoreBookNumberRange = 10
	client                        = &http.Client{}
	datafile                      = "./data"
	apiURL                        = "https://xkcd.com/%v/info.0.json"
)

// Commic is Commic.
type Commic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safeTitle"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// Commics is Commics
type Commics struct {
	CommicsList []Commic `json:"commics"`
}

func l(vs ...interface{}) {
	for _, v := range vs {
		fmt.Println(v)
	}
}

func usageAndExit() {
	l("")
	l(" Initialize:\n  $ go run main.go init\n")
	l(" Search:\n  $ go run main.go search <book title>")
	l("")
	os.Exit(1)
}

func get(url string) (*Commic, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		l(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		l(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l(err)
		return nil, err
	}

	var commic Commic
	json.Unmarshal(body, &commic)
	return &commic, nil
}

func store(commics *Commics) error {
	bytes, err := json.Marshal(commics)
	if err != nil {
		panic(err)
	}
	writeFile(datafile, bytes)
	return nil
}

func writeFile(filename string, bytes []byte) {
	ioutil.WriteFile(filename, bytes, os.ModePerm)
}

func initDB() {
	commicsList := []Commic{}
	for i := minGetAndStoreBookNumberRange; i < maxGetAndStoreBookNumberRange; i++ {
		commic, err := get(fmt.Sprintf(apiURL, i))
		if err != nil {
			panic(err)
		}
		commicsList = append(commicsList, *commic)
	}
	commics := Commics{}
	commics.CommicsList = commicsList
	err := store(&commics)
	if err != nil {
		panic(err)
	}
}

func search(key string) {
	bytes, err := ioutil.ReadFile(datafile)
	if err != nil {
		panic(err)
	}
	var commics Commics
	if err != json.Unmarshal(bytes, &commics) {
		panic(err)
	}
	for _, c := range commics.CommicsList {
		if c.Title == key {
			l("=== We found book! ===")
			l("URL:", c.Link)
			l("Transcript:", c.Transcript)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		usageAndExit()
	}
	command := os.Args[1]
	switch command {
	case "init":
		initDB()
	case "search":
		if len(os.Args) < 3 {
			usageAndExit()
		}
		keyword := os.Args[2]
		search(keyword)
	default:
		usageAndExit()
	}
}
