package main

import (
	"fmt"
	"os"
	"strings"
)

// 字符串拼接，go run strings/join.go a b c
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
