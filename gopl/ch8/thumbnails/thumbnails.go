package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	filenames := make(chan string, 10)
	for i := 0; i < 10; i++ {
		filenames <- "Msg" + strconv.FormatInt(int64(i), 10)
	}
	close(filenames)

	size := makeThumbnails6(filenames)
	fmt.Println("size=" + strconv.FormatInt(int64(size), 10))
}

func makeThumbnails6(filenames <-chan string) int {
	sizes := make(chan int)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			sizes <- len(thumb)
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int
	for size := range sizes {
		total += size
	}
	return total
}

func ImageFile(f string) (string, error) {
	time.Sleep(1 * time.Second)
	return f, nil
}
