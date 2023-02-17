package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "飞机666"
	fmt.Println(len(s))                    // 字节数：9
	fmt.Println(utf8.RuneCountInString(s)) // 字符数：5

	// 索引操作，取的是字节，uint8 类型
	u := s[0]                          // uint8 类型，值为 233
	u2 := s[7]                         // uint8 类型，值为 54
	fmt.Println(string(u), string(u2)) // "é" "6"

	// 子串操作，也是按字节取的
	s2 := s[0:1] // �
	s3 := s[0:2] // ��
	s4 := s[0:3] // 飞
	s5 := s[0:4] // 飞�
	fmt.Println(s2, s3, s4, s5)

	// 字符串不可修改
	//s[7] = 7 // 编译报错：Cannot assign to s[7]

	/* 转义字符
	\a      响铃
	\b      退格
	\f      换页
	\n      换行
	\r      回车
	\t      制表符
	\v      垂直制表符
	\'      单引号（只用在 '\'' 形式的rune符号面值中）
	\"      双引号（只用在 "..." 形式的字符串面值中）
	\\      反斜杠
	*/

	// 转义字符串
	abcX := "\x41\x42\x43" // 十六进制转义字符串："ABC"，仅支持单字节转义
	abc0 := "\101\102\103" // 八进制转义字符串："ABC"，仅支持单字节转义
	中国u := "\u4e2d\u56fd"  // 十六进制转义字符串："中国"，支持多字节转义
	fmt.Println(abcX, abc0, 中国u)

	// 原生字符串面值
	a := `abc\tdef` // 原生字符串面值不转义
	b := "abc\tdef" // 非原生字符串面值会转移
	fmt.Println(a, b)

	// 原生字符串面值可跨域多行
	const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`
	fmt.Println(GoUsage)
}
