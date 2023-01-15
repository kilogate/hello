package main

import "fmt"

func main() {
	testMake()
}

func testMake() {
	// 1 数组不用make可以直接使用
	var strArray [1]string   // 长度为1元素为空的数组
	fmt.Println(strArray[0]) // 可读
	strArray[0] = "ABC"      // 可写

	// 2 切片需要make或append后再使用
	var strSlice []string // nil
	//fmt.Println(strSlice[0])  // 读报错，runtime error: index out of range [0] with length 0
	//strSlice[0] = "ABC" // 写报错，runtime error: index out of range [0] with length 0

	// nil切片支持直接append，append之后返回新切片
	strSlice = append(strSlice, "a") // 长度为1容量为1元素为a的切片
	fmt.Println(strSlice[0])         // 下标大于等于长度的时候会报错，runtime error: index out of range [0] with length 0
	strSlice[0] = "ABC"              // 下标大于等于长度的时候会报错，runtime error: index out of range [0] with length 0

	// make创建指定长度和容量的切片
	strSlice = make([]string, 3, 10) // 长度为3容量为10元素为空的切片
	fmt.Println(strSlice[0])         // 下标大于等于长度的时候会报错，runtime error: index out of range [0] with length 0
	strSlice[0] = "ABC"              // 下标大于等于长度的时候会报错，runtime error: index out of range [0] with length 0

	// 3 map需要make后再使用
	var strMap map[string]string // nil
	s, ok := strMap["a"]         // 读不到值，不会报错，返回："", false
	fmt.Println(s, ok)
	//strMap["k"] = "v"            // 写报错，assignment to entry in nil map

	// make创建指定容量的map
	strMap = make(map[string]string, 10) // 容量为10元素为空的map
	fmt.Println(strMap["k"])
	strMap["k"] = "v"

	// 4 chan需要make后再使用
	var strChan chan string // nil
	strChan = make(chan string)
	go func() {
		strChan <- "abc"
	}()
	fmt.Println(<-strChan)
}
