package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

var done = make(chan struct{})

func fetch(url string) (filename string, n int64, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go func(u string) {
			local, n, err := fetch(u)
			if err == nil {
				fmt.Println("end : ", u, local, n)
				close(done)
			} else {
				fmt.Println("canceled : ", u, local, n)
			}
			wg.Done()
		}(url)
	}
	wg.Wait()
}
