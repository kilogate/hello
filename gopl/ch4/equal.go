package main

import "fmt"

func main() {
	x := []string{"A", "B", "C"}
	y := []string{"A", "B", "C"}
	//fmt.Println(x == y) // // 切片类型不支持 == 比较，编译报错
	fmt.Println(equal(x, y))

	a1 := [3]string{"A", "B", "C"}
	a2 := [3]string{"A", "B", "C"}
	fmt.Println(a1 == a2) // 数组可以直接比较

	a3 := [1]interface{}{x}
	a4 := [1]interface{}{y}
	fmt.Println(a3 == a4) // panic: runtime error: comparing uncomparable type []string
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
