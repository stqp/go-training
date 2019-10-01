package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ifnotExistsThenCreateDir(path string) {
	path = filepath.Dir(path)
	if len(path) <= 0 {
		return
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0700); err != nil {
			panic(fmt.Sprintf("failed to create output directory : %s", path))
		}
	}
}

func absoluteOutDir() string {
	cur, _ := os.Getwd()
	return cur + "/" + out
}
