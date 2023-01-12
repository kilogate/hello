package main

import (
	"fmt"
	"strings"
)

func main() {
	testSplitAndJoin()
	testHasPrefixAndHasSuffix()
}

func testSplitAndJoin() {
	// 字符串分割
	split := strings.Split("a,b,c", ",")
	fmt.Println(split)

	// 字符串拼接
	join := strings.Join([]string{"1", "2", "3"}, " ")
	fmt.Println(join)
}

func testHasPrefixAndHasSuffix() {
	// 前缀匹配
	fmt.Println(strings.HasPrefix("abc", "a"))

	// 后缀匹配
	fmt.Println(strings.HasSuffix("abc", "c"))
}
