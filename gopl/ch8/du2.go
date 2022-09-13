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

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 8)

func main() {
	now := time.Now()

	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	go func() {
		for _, root := range roots {
			wg.Add(1)
			go walkDir2(root, &wg, fileSizes)
		}
	}()
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	tick := time.Tick(100 * time.Millisecond)

	var nFiles, nBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nFiles++
			nBytes += size
		case <-tick:
			printDiskUsage2(nFiles, nBytes)
		}
	}
	printDiskUsage2(nFiles, nBytes) // final totals
	fmt.Println()
	fmt.Printf("cost: %fs\n", time.Since(now).Seconds())
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir2(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirEntries2(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir2(subDir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirEntries returns the entries of directory dir.
func dirEntries2(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage2(nFiles, nBytes int64) {
	fmt.Printf("\r%d files  %.1f GB", nFiles, float64(nBytes)/1e9)
}
