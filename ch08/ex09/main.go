package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type FileSize struct {
	size int64
	root string
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan FileSize)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(1000 * time.Millisecond)
	}

	fileSizeMap := make(map[string]int64)
loop:
	for {
		select {
		case f, ok := <-fileSizes:
			if !ok {
				break loop
			}
			fileSizeMap[f.root] += f.size
		case <-tick:
			printDiskUsage(&fileSizeMap)
		}
	}

	printDiskUsage(&fileSizeMap)
}

func printDiskUsage(m *map[string]int64) {
	for root, size := range *m {
		fmt.Printf("root : %s , %.1f GB\n", root, float64(size)/1e9)
	}
	fmt.Println("")
}

func walkDir(root string, dir string, n *sync.WaitGroup, fileSizes chan<- FileSize) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, fileSizes)
		} else {
			fileSizes <- FileSize{size: entry.Size(), root: root}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
