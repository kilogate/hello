package main

import (
	"fmt"
)

func main() {
	testSwitchFallthrough()
	testSwitchTrue()
}

func testSwitchFallthrough() {
	str := "C"
	switch str {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")
		fallthrough
	case "D":
		fmt.Println("D")
		fallthrough
	case "E":
		fmt.Println("E")
	case "F":
		fmt.Println("F")
	}
}

func testSwitchTrue() {
	switch { // 等价于 switch true，可用于多个 if else，这种形式叫做无 tag switch，即 tagless switch
	case 4 > 5:
		fmt.Println("4 > 5")
	case 5 > 4:
		fmt.Println("5 > 4")
	default:
		fmt.Println("default")
	}
}
