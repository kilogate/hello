package main

import (
	"fmt"
	"strconv"
)

func main() {
	testItoaAndAtoi()
	testFormat()
	testParse()
}

func testItoaAndAtoi() {
	// int -> string
	a := strconv.Itoa(123)
	fmt.Println(a)

	// string -> int
	i, err := strconv.Atoi(a)
	fmt.Println(i, err)
}

func testFormat() {
	// int64 -> string
	base2 := strconv.FormatInt(11, 2)   // 二进制
	base8 := strconv.FormatInt(11, 8)   // 八进制
	base10 := strconv.FormatInt(11, 10) // 十进制
	base16 := strconv.FormatInt(11, 16) // 十六进制
	fmt.Println(base2, base8, base10, base16)

	// uint64 -> string
	base2 = strconv.FormatUint(11, 2)   // 二进制
	base8 = strconv.FormatUint(11, 8)   // 八进制
	base10 = strconv.FormatUint(11, 10) // 十进制
	base16 = strconv.FormatUint(11, 16) // 十六进制
	fmt.Println(base2, base8, base10, base16)
}

func testParse() {
	// string -> int64
	i127, err := strconv.ParseInt("127", 10, 8) // bitSize为8表示转换最大值不超过8位比特，即最大127，一般bitSize传64即可
	i128, err := strconv.ParseInt("128", 10, 8) // 超过最大值时返回的是最大值127，且此处会报错（err非nil）
	fmt.Println(i127, i128, err)

	// string -> uint64
	u255, err := strconv.ParseUint("255", 10, 8) // 无符号数最大255，返回类型：uint64
	u256, err := strconv.ParseUint("256", 10, 8) // 超过最大值时返回的是最大值255，且此处会报错（err非nil）
	fmt.Println(u255, u256, err)
}
