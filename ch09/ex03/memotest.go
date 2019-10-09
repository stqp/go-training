package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string, done chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	time.Sleep(1000 * time.Millisecond) // テストのために
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

type NumberedUrl struct {
	i   int
	url string
}

func incomingURLs() <-chan NumberedUrl {
	ch := make(chan NumberedUrl)
	go func() {
		for i, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- NumberedUrl{i: i, url: url}
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done chan struct{}) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for nu := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(nu.url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			nu.url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup

	for nu := range incomingURLs() {
		n.Add(1)
		go func(nu NumberedUrl) {
			defer n.Done()
			start := time.Now()
			done := make(chan struct{})

			// 適当にこの辺の番号だったらリクエストをキャンセルする。
			if nu.i == 4 || nu.i == 5 {
				go func(done *chan struct{}) {
					close(*done)
				}(&done)
			}

			value, err := m.Get(nu.url, done)
			if err != nil {
				//fmt.Println(err)
				//return
			}
			if value != nil {
				fmt.Printf("%s, %s, %d bytes\n",
					nu.url, time.Since(start), len(value.([]byte)))
			} else {
				fmt.Printf("%s, %s, %d bytes\n",
					nu.url, time.Since(start), 0)
			}

		}(nu)
	}
	n.Wait()
}
