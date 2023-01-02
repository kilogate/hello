package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串分割
	split := strings.Split("a,b,c", ",")
	fmt.Println(split)

	// 字符串拼接
	join := strings.Join([]string{"1", "2", "3"}, " ")
	fmt.Println(join)
}
