package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	testUTF8()
}

func testUTF8() {
	s := "飞机666"
	fmt.Println(len(s))                    // 字节数：9
	fmt.Println(utf8.RuneCountInString(s)) // 字符数：5
}
