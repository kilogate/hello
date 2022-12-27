package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 标准输入，直接运行，按行输入，通过 Ctrl + d 结束输入
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
