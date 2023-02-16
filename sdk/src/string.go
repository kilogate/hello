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
	u := s[0]                          // uint8 类型，值为 233
	u2 := s[7]                         // uint8 类型，值为 54
	fmt.Println(string(u), string(u2)) // "é" "6"

	// 取子串也是按字节取的
	s2 := s[0:1] // �
	s3 := s[0:2] // ��
	s4 := s[0:3] // 飞
	s5 := s[0:4] // 飞�
	fmt.Println(s2, s3, s4, s5)

	// 字符串不可修改
	//s[7] = 7 // 编译报错：Cannot assign to s[7]
}
