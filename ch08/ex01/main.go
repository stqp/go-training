package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		//time.Sleep(1 * time.Second)
		c.Close()
	}
}

func main() {
	portNum := flag.Int("port", 8000, "")
	flag.Parse()
	port := strconv.Itoa(*portNum)
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
