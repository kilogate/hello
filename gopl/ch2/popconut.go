package main

import "fmt"

// pc[i] is the population count of 1
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 匿名函数初始化
//var pc = func() (pc [256]byte) {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//	return
//}()

func PopCount(x uint64) int {
	count := pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]
	return int(count)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(PopCount(uint64(i)))
	}
}
