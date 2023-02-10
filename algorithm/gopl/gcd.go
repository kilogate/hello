package main

import "fmt"

func main() {
	fmt.Println(gcd(12, 18))
	fmt.Println(gcd(18, 12))
	fmt.Println(gcd(24, 12))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}
