package main

import (
	"log"
	"time"

	"github.com/samber/lo"
)

func main() {
	// Synchronize
	lo.Synchronize().Do(func() {
		log.Println("async task start")
		time.Sleep(time.Second * 3)
		log.Println("async task end")
	})

	// Async1
	async1 := lo.Async1(func() string {
		log.Println("Async1 start")
		time.Sleep(time.Second * 3)
		log.Println("Async1 end")
		return "Async1 done"
	})

	// Async2
	async2 := lo.Async2(func() (string, string) {
		log.Println("Async2 start")
		time.Sleep(time.Second * 2)
		log.Println("Async2 end")
		return "Async1 done1", "Async1 done2"
	})

	// 等待 Async1 执行完成
	log.Printf("async1: %v\n", <-async1)
	// 等待 Async2 执行完成
	res := <-async2
	log.Printf("async2: %v\n", res)
}
