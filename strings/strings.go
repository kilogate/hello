package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 字符串拼接，go run strings/strings.go a b c
	fmt.Println(strings.Join(os.Args[1:], " "))
}
