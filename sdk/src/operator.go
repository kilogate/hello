package main

import "fmt"

func main() {
	testMod()
}

func testMod() {
	// 取模运算符的符号和被取模数的符号总是一致的
	fmt.Println(5 % 3)   // 2
	fmt.Println(5 % -3)  // 2
	fmt.Println(-5 % 3)  // -2
	fmt.Println(-5 % -3) // -2

	// 整数除法会截断余数
	fmt.Println(9 / 5.0) // 1.8，浮点数除法
	fmt.Println(9 / 5)   // 1，整数除法
}
