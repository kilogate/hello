package main

import (
	"fmt"
	"os"
)

func main() {
	// 命令行参数，go run os/os.go a b c
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
