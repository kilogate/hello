package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")

	// 取消命令
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	// 倒计时
	tick := time.Tick(1 * time.Second)

	// 监听
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Aborted.")
			return
		}
	}
	fmt.Println("launch...")
}
