package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("pass your name.")
		return
	}
	name := os.Args[1]

	origin := "http://localhost:8000"
	u := "ws://localhost:8000/?name=" + name
	conn, err := websocket.Dial(u, "", origin)
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
