package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "飞机666"
	fmt.Println(len(s))                    // 字节数：9
	fmt.Println(utf8.RuneCountInString(s)) // 字符数：5

	// 取的是字节，uint8 类型
	u := s[0]
	fmt.Println(string(u))
	u2 := s[7]
	fmt.Println(string(u2))
}
