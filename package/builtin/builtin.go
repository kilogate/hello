package main

import "fmt"

func main() {
	testCollection()
}

func testCollection() {
	// 数组不用make可直接使用
	var strArray [1]string
	strArray[0] = "ABC"
	fmt.Println(strArray)

	// 切片不用make也可以直接使用
	var strSlice []string
	strSlice = append(strSlice, "a")
	fmt.Println(strSlice)

	// map必须make，否则使用时会 panic: assignment to entry in nil map
	var strMap map[string]string
	strMap = make(map[string]string)
	strMap["k"] = "v"
	fmt.Println(strMap)
}
