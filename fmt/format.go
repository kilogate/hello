package main

import (
	"fmt"
	"math"
)

// 格式化动词和转义字符
func main() {
	// %x, %d, %o, %b ：十六进制整数，十进制整数，八进制整数，二进制整数
	fmt.Printf("%x, %d, %o, %b\n", 13, 13, 13, 13)

	// %f, %g, %e：浮点数3.141593，浮点数3.141592653589793，浮点数3.141593e+00
	fmt.Printf("%f, %g, %e\n", math.Pi, math.Pi, math.Pi)

	// %t：布尔值，true或false
	fmt.Printf("%t, %t\n", true, false)

	// %c：字符（rune） (Unicode码点)
	fmt.Printf("%c, %c, %c\n", 'a', 'b', 'c')

	// %s：字符串
	fmt.Printf("%s\n", "字符串")

	// %q：带双引号的字符串"abc"或带单引号的字符'd'
	fmt.Printf("带双引号的字符串%q或带单引号的字符%q\n", "abc", 'd')

	// %v：变量的自然形式（natural format）
	fmt.Printf("%v, %v, %v, %v, %v\n", 13, math.Pi, true, 'a', "abc")

	// %T：变量的类型
	fmt.Printf("%T, %T, %T, %T, %T\n", 13, math.Pi, true, 'a', "abc")

	// %%：百分号
	fmt.Printf("%%\n")

	// \n：换行符、\t：制表符
	fmt.Printf("a\tb\n")
}
