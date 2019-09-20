package main

import (
	"io"
)

func main() {}

type cw struct {
	count  *int64
	writer io.Writer
}

func (w cw) Write(p []byte) (n int, err error) {
	*w.count += int64(len(p))
	return w.writer.Write(p)
}

// CountingWriter counts bytes
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	if w == nil {
		return nil, nil
	}

	count := int64(0)
	newWriter := cw{
		writer: w,
		count:  &count,
	}

	return newWriter, newWriter.count
}
