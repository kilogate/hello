package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")

	abort := make(chan string)
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			abort <- input.Text()
		}
	}()

	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-abort:
			fmt.Println("aborted.")
			return
		case <-tick:
			fmt.Println(countdown)
		}
	}
	fmt.Println("launch...")
}
