package main

import (
	"fmt"
	"sync"
)

var balanceMutex sync.Mutex
var balance int

func Deposit(amount int) {
	balanceMutex.Lock()
	defer balanceMutex.Unlock()

	balance = balance + amount
}

func Balance() int {
	balanceMutex.Lock()
	defer balanceMutex.Unlock()

	return balance
}

func main() {
	fmt.Printf("main start, balance: %v\n", Balance())
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				Deposit(1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("main end, balance: %v\n", Balance())
}
