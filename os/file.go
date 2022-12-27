package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1 从标准输入中读取：直接运行，按行输入，通过 Ctrl + d 结束输入
// 2 从文件中读取：go run os/file.go go.mod os/file.go
func main() {
	counts := make(map[string]int)
	paths := os.Args[1:]
	if len(paths) == 0 { // 1 从标准输入中读取
		countLines(os.Stdin, counts)
	} else { // 2 从文件中读取
		for _, path := range paths {
			file, err := os.Open(path) // 打开文件
			if err != nil {
				fmt.Fprintf(os.Stderr, "process fail: %v\n", err) // 错误信息打印到标准错误流
				continue
			}
			countLines(file, counts)
			file.Close() // 关闭文件
		}
	}

	// 打印重复行次数
	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}
