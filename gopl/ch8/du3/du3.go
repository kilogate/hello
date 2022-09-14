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

var done = make(chan struct{})

// main go run gopl/ch8/du3/du3.go ~/
func main() {
	now := time.Now()

	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Traverse the file tree.
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	go func() {
		for _, root := range roots {
			wg.Add(1)
			go walkDir(root, &wg, fileSizes)
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
			printDiskUsage(nFiles, nBytes)
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing.
			}
			fmt.Println("任务被停止")
			return
		}
	}
	printDiskUsage(nFiles, nBytes) // final totals
	fmt.Println()
	fmt.Printf("cost: %fs\n", time.Since(now).Seconds())
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subDir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirEntries returns the entries of directory dir.
func dirEntries(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}

	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("\r%d files  %.1f GB", nFiles, float64(nBytes)/1e9)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
