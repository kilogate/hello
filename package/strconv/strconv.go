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
	a := strconv.Itoa(123)
	i, err := strconv.Atoi(a)
	fmt.Println(i, err)
}

func testFormat() {
	base2 := strconv.FormatInt(11, 2)
	base8 := strconv.FormatInt(11, 8)
	base10 := strconv.FormatInt(11, 10)
	base16 := strconv.FormatInt(11, 16)
	fmt.Println(base2, base8, base10, base16)

	base2 = strconv.FormatUint(11, 2)
	base8 = strconv.FormatUint(11, 8)
	base10 = strconv.FormatUint(11, 10)
	base16 = strconv.FormatUint(11, 16)
	fmt.Println(base2, base8, base10, base16)
}

func testParse() {
	// 一般bitSize传64即可
	i127, err := strconv.ParseInt("127", 10, 8) // bitSize为8表示转换最大值不超过8位比特，即最大127
	fmt.Println(i127, err)

	i128, err := strconv.ParseInt("128", 10, 8) // 超过最大值时返回的是最大值127，但err非nil
	fmt.Println(i128, err)

	u128, err := strconv.ParseUint("128", 10, 8) // 无符号数最大255
	fmt.Println(u128, err)

	u255, err := strconv.ParseUint("255", 10, 8) // 无符号数最大255
	fmt.Println(u255, err)

	u256, err := strconv.ParseUint("256", 10, 8) // 无符号数最大255，此处会报错
	fmt.Println(u256, err)
}
