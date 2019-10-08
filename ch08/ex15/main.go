package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	updated time.Time
	ch      *chan string
	conn    *net.Conn
	who     string
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan string)
)

var clients map[*client]bool

func broadcaster() {
	clients = make(map[*client]bool)
	for {
		select {
		case msg := <-messages:
			for c := range clients {
				*c.ch <- msg
				fmt.Println(c.who)
			}

		case c := <-entering:
			clients[c] = true

		case c := <-leaving:
			close(*c.ch)
			(*c.conn).Close()
			delete(clients, c)
		}
	}
}

func handleConn(conn net.Conn) {

	who := conn.RemoteAddr().String()
	ch := make(chan string, 10)
	c := &client{
		updated: time.Now(),
		ch:      &ch,
		conn:    &conn,
		who:     who,
	}
	go clientWriter(conn, *c.ch)

	*c.ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- c
	input := bufio.NewScanner(conn)
	for input.Scan() {
		c.updated = time.Now()
		messages <- who + ": " + input.Text()
	}

	_, ok := <-*c.ch

	if ok {
		leaving <- c
		messages <- who + " has left"
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func closeIfNoActionFor(sec int) {
	for {
		for c, exists := range clients {
			if exists && time.Now().Sub(c.updated).Seconds() > float64(sec) {
				leaving <- c
				messages <- c.who + " has left"
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	go closeIfNoActionFor(600) // 5 minutes.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
