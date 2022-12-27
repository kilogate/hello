package main

import (
	"bufio"
	"fmt"
	"os"
)

// 标准输入，直接运行，按行输入，通过 Ctrl + d 结束输入
func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
