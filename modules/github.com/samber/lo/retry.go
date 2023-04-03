package main

import (
	"log"
	"time"

	"github.com/samber/lo"
)

func main() {
	//testDebounce()
	//testDebounceBy()
}

func testDebounce() {
	// NewDebounce
	reset, cancel := lo.NewDebounce(time.Second, func() { log.Println("执行任务一") }, func() { log.Println("执行任务二") })
	reset()  // 执行一次
	cancel() // 取消执行

	// 等待异步执行结束
	time.Sleep(time.Minute)
}

func testDebounceBy() {
	// NewDebounceBy
	reset, cancel := lo.NewDebounceBy(time.Second, func(key string, count int) { log.Printf("执行任务, key: %s, count: %d\n", key, count) })
	reset("samuel")
	reset("samuel")
	reset("john")
	reset("samuel")
	reset("samuel")
	time.Sleep(500 * time.Millisecond)
	cancel("john")

	// 等待异步执行结束
	time.Sleep(time.Minute)
}
