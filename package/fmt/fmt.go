package main

import (
	"fmt"
	"math"
)

type Stu struct {
	Name string
	Age  int
}

func main() {
	testPrintf()
	testSprintf()
}

// testPrintf 格式化动词和转义字符
func testPrintf() {
	// %x, %d, %o, %b ：十六进制整数，十进制整数，八进制整数，二进制整数
	fmt.Printf("%x, %d, %o, %b\n", 13, 13, 13, 13)

	fmt.Printf("%+d\n", 13) // 带符号
	fmt.Printf("%#o\n", 13) // 带零前导
	fmt.Printf("%X\n", 13)  // 大写十六进制
	fmt.Printf("%#x\n", 13) // 带0x

	fmt.Printf("%5d\n", 13)  // 长度5，右对齐
	fmt.Printf("%-5d\n", 13) // 长度5，左对齐
	fmt.Printf("%05d\n", 13) // 长度5，左边补零

	// %f, %g, %e：浮点数3.141593，浮点数3.141592653589793，浮点数3.141593e+00
	fmt.Printf("%f, %g, %e\n", math.Pi, math.Pi, math.Pi)

	fmt.Printf("%.2f\n", 111.222)    // 两位小数
	fmt.Printf("%4.2f\n", 111.222)   // 两位小数，最小宽度4
	fmt.Printf("%010.2f\n", 111.222) // 两位小数，最小宽度10，不足补零

	// %t：布尔值，true或false
	fmt.Printf("%t, %t\n", true, false)

	// %c：字符（rune） (Unicode码点)
	fmt.Printf("%c, %c, %c\n", 'a', 'b', 'c')

	// %s：字符串
	fmt.Printf("%s\n", "字符串")
	fmt.Printf("%5s\n", "字符串")         // 展示的最小宽度为5，右对齐
	fmt.Printf("%-5s\n", "字符串")        // 左对齐
	fmt.Printf("%05s\n", "字符串")        // 补零
	fmt.Printf("%.5s\n", "123456789")  // 截断的最大宽度5
	fmt.Printf("%5.7s\n", "123456789") // 展示的最小宽度为5，截断的最大宽度7
	fmt.Printf("%5.3s\n", "123456789") // 展示的最小宽度为5，截断的最大宽度3

	// %q：带双引号的字符串"abc"或带单引号的字符'd'
	fmt.Printf("带双引号的字符串%q或带单引号的字符%q\n", "abc", 'd')
	fmt.Printf("%q\n", `a"b"c`)  // 引号被转义
	fmt.Printf("%#q\n", `a"b"c`) // 反引号

	// %U：Unicode
	fmt.Printf("%d %X\n", '中', '中')
	fmt.Printf("%U\n", '中')
	fmt.Printf("%#U\n", '中') // 带字符

	// %v：变量的自然形式（natural format），打印结构体使用
	fmt.Printf("%v, %v, %v, %v, %v\n", 13, math.Pi, true, 'a', "abc")

	stu := &Stu{"Lask", 30}
	fmt.Printf("%v\n", stu)  // 仅打印值
	fmt.Printf("%+v\n", stu) // 再加上字段名
	fmt.Printf("%#v\n", stu) // 再加上报名和类型名

	// %p：指针
	fmt.Printf("%p\n", stu)
	fmt.Printf("%#p\n", stu) // 不带0x

	// %T：变量的类型
	fmt.Printf("%T, %T, %T, %T, %T\n", 13, math.Pi, true, 'a', "abc")

	// %%：百分号
	fmt.Printf("%%\n")

	// \n：换行符、\t：制表符
	fmt.Printf("a\tb\n")
}

func testSprintf() {
	// 格式化并返回字符串
	str := fmt.Sprintf("%s", "abc")
	fmt.Println(str)
}
