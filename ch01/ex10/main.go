package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	urlTool "net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	fmt.Println(time.Now().Format("15:04:05"))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	u, err := urlTool.Parse(url)
	if err != nil {
		panic(err)
	}
	outputFile, err := os.Create("ex10/out/" + u.Host + "_" + time.Now().Format("15:04:05"))
	if err != nil {
		log.Fatal(err)
	}

	nbytes, err := io.Copy(outputFile, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
