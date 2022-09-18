package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func main() {
	urls := []string{
		"https://baidu.com",
		"https://alibaba.com",
		"https://tencent.com",
		"https://58.com",
		"https://bytedance.com",
	}

	memo := New(httpGetBody)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()

				start := time.Now()
				value, err := memo.Get(url)
				if err != nil {
					log.Print(err)
				}
				fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			}(url)
		}
	}

	wg.Wait()
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
