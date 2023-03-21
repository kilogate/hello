package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	//testScanner()
	testReader()
}

func testScanner() {
	// 从标准输入（控制台）中输入，按行输入，通过 Ctrl + d 结束输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

func testReader() {
	// 从控制台输入，通过 Ctrl + d 退出输入
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("ReadRune occurs err, err: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			fmt.Printf("ReadRune found invalid rune, r: %q\n", r)
			continue
		}
		fmt.Printf("rune: %q, bytes: %d\n", r, n)
	}
}
