package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	ch   chan string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			msg := "= member list =\n"
			for c := range clients {
				msg += fmt.Sprintf("・%s\n", c.name)
			}
			cli.ch <- msg

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	c := client{
		ch: make(chan string),
	}
	go clientWriter(conn, c.ch)

	who := conn.RemoteAddr().String()
	c.name = who
	c.ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- c

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- c
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
