package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func parseArg(args []string) map[string]string {
	m := make(map[string]string)
	for _, arg := range args {
		kv := strings.Split(arg, "=")
		k := kv[0]
		v := kv[1]
		m[k] = v
	}
	return m
}

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("no args. end.")
		os.Exit(1)
	}
	m := parseArg(os.Args[1:])
	for {
		t := ""
		for k, v := range m {
			t += fmt.Sprintf("%s: %s", k, getTime(v))
		}
		fmt.Println(t)
		time.Sleep(1 * time.Second)
	}
}

func getTime(host string) string {
	buffer := &bytes.Buffer{}
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(buffer, conn)
	return buffer.String()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
