package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defaultCount := 5

	input := bufio.NewScanner(c)
	count := defaultCount
	go func() {
		for ; count > 0; count-- {
			select {
			case <-time.Tick(1 * time.Second):
				fmt.Println(count)
			}
		}
		c.Close()
	}()

	for input.Scan() {
		count = defaultCount
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
