package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var idx uint64 = 0
var tokens = make(chan struct{}, 20)

// main go run gopl/ch8/crawl/crawl.go www.baidu.com www.ali.com www.tencent.com
func main() {
	workList := make(chan []string)
	var n = 0

	// Start with the command-line arguments.
	n++
	go func() { workList <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} // acquire a token
	list := extract(url)
	<-tokens // release the token

	return list
}

func extract(url string) []string {
	if len(url) > 50 {
		return nil
	}

	var list []string
	for i := 0; i < 2; i++ {
		list = append(list, fmt.Sprintf("[%v]%s_%v", atomic.AddUint64(&idx, 1), url, i))
	}
	time.Sleep(time.Second)
	return list
}
