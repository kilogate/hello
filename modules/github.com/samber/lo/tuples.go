package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	t2 := lo.T2("Tom", 26)
	fmt.Println(t2)

	a, b := lo.Unpack2(t2)
	fmt.Println(a, b)

	zip2 := lo.Zip2([]string{"A", "B", "C"}, []int{123, 456})
	fmt.Println(zip2)

	a1, b1 := lo.Unzip2(zip2)
	fmt.Println(a1, b1)
}
