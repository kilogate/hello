package main

import (
	"fmt"
	"os"
)

// 命令行参数，go run os/args.go a b c
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
