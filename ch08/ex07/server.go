package main

import (
	"fmt"
	"net/http"
)

func serve() {
	fmt.Println(http.Dir(absoluteOutDir()))
	fs := http.FileServer(http.Dir(absoluteOutDir()))
	http.Handle("/", fs)
	http.ListenAndServe(hostname+":"+port, nil)
}
