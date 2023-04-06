package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	// Range
	fmt.Println(lo.Range(5))  // [0 1 2 3 4]
	fmt.Println(lo.Range(-5)) // [0 -1 -2 -3 -4]

	// RangeFrom
	fmt.Println(lo.RangeFrom(3, 5))  // [3 4 5 6 7]
	fmt.Println(lo.RangeFrom(3, -5)) // [3 2 1 0 -1]

	// RangeWithSteps
	fmt.Println(lo.RangeWithSteps(3, 10, 3))   // [3 6 9]
	fmt.Println(lo.RangeWithSteps(-10, -3, 3)) // [-10 -7 -4]

	// Clamp
	fmt.Println(lo.Clamp(-5, 1, 10)) // 1
	fmt.Println(lo.Clamp(5, 1, 10))  // 5
	fmt.Println(lo.Clamp(50, 1, 10)) // 10

	// Sum & SumBy：求和
	fmt.Println(lo.Sum([]int{1, 3, 5, 7}))                                           // 16
	fmt.Println(lo.SumBy([]int{1, 3, 5, 7}, func(item int) int { return item + 1 })) // 20
}
