package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var imgType string

func main() {
	flag.StringVar(&imgType, "t", "", "")
	flag.Parse()
	if err := toImage(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toImage(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch imgType {
	case "png":
		return png.Encode(out, img)
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 255})
	case "":
		fmt.Println("specify output image type.")
		os.Exit(1)
	default:
		fmt.Println("specified output image type is not supported.")
		os.Exit(1)
	}
	return nil
}
