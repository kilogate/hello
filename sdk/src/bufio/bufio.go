package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//testScanner()
}

func testScanner() {
	// 从标准输入（控制台）中输入，按行输入，通过 Ctrl + d 结束输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}
