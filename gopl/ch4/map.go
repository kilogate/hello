package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
	}

	// 无序打印
	for i := 0; i < 10; i++ {
		for k, v := range m {
			fmt.Print(k, v, ",")
		}
		fmt.Println()
	}

	fmt.Println()

	// 有序打印
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	for i := 0; i < 10; i++ {
		for _, k := range ks {
			fmt.Print(k, m[k], ",")
		}
		fmt.Println()
	}
}
