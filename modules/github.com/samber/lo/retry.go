package main

import (
	"fmt"
	"log"
	"time"

	"github.com/samber/lo"
)

func main() {
	//testDebounce()
	//testDebounceBy()
	//testAttempt()
	//testAttemptWithDelay()
	//testAttemptWhile()
	//testAttemptWhileWithDelay()
	testTransaction()
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

func testAttempt() {
	maxIteration, err := lo.Attempt(3, func(index int) error {
		log.Printf("第 %d 次尝试执行", index+1)
		if index >= 1 {
			return nil
		}
		return fmt.Errorf("执行失败")
	})
	log.Printf("第 %d 次执行成功，报错信息: %v", maxIteration, err)
}

func testAttemptWithDelay() {
	maxIteration, since, err := lo.AttemptWithDelay(5, time.Second, func(index int, duration time.Duration) error {
		log.Printf("第 %d 次尝试执行, 已过去 %v", index+1, duration)
		if index >= 3 {
			return nil
		}
		return fmt.Errorf("执行失败")
	})
	fmt.Printf("第 %d 次执行成功，已过去 %v, 报错信息: %v", maxIteration, since, err)
}

func testAttemptWhile() {
	maxIteration, err := lo.AttemptWhile(5, func(index int) (error, bool) {
		log.Printf("第 %d 次尝试执行", index+1)
		if index >= 3 {
			return nil, true
		}
		return fmt.Errorf("执行失败"), true
	})
	fmt.Printf("第 %d 次执行成功，报错信息: %v", maxIteration, err)
}

func testAttemptWhileWithDelay() {
	maxIteration, since, err := lo.AttemptWhileWithDelay(5, time.Second, func(index int, duration time.Duration) (error, bool) {
		log.Printf("第 %d 次尝试执行, 已过去 %v", index+1, duration)
		if index >= 3 {
			return nil, true
		}
		return fmt.Errorf("执行失败"), true
	})
	fmt.Printf("第 %d 次执行成功，已过去 %v, 报错信息: %v", maxIteration, since, err)
}

func testTransaction() {
	exec := func(state int) (int, error) {
		if state == 3 {
			return 3, fmt.Errorf("执行失败")
		}

		res := state + 1
		log.Printf("exec start, state: %d -> %d\n", state, res)
		return res, nil
	}
	rollback := func(state int) int {
		res := state - 1
		log.Printf("rollback start, state: %d -> %d\n", state, res)
		return res
	}
	res, err := lo.NewTransaction[int]().
		Then(exec, rollback).
		Then(exec, rollback).
		Then(exec, rollback).
		Process(1)
	log.Printf("执行结果: %d, 报错信息: %v", res, err)
}
