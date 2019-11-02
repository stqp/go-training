package main

import (
	"archive/tar"
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var arcType string
var arcFilePath string

func main() {
	flag.StringVar(&arcType, "t", "", "")
	flag.StringVar(&arcFilePath, "f", "", "")
	flag.Parse()

	if arcType == "" || arcFilePath == "" {
		fmt.Println("specify archive file type , and archive file path.")
		os.Exit(1)
	}

	if err := readArchive(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func readArchive() error {

	f, err := os.Open(arcFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fs, err := f.Stat()

	switch arcType {
	case "zip":
		r, err := zip.NewReader(f, fs.Size())
		if err != nil {
			fmt.Println(err)
		}
		for _, zf := range r.File {
			fmt.Printf("zip: %s\n", zf.Name)
		}
	case "tar":
		r := tar.NewReader(f)
		if err != nil {
			fmt.Println(err)
		}
		for {
			hdr, err := r.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("tar: %s\n", hdr.Name)
		}
	}

	return nil
}
