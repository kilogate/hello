package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	testArgs()
	testOpen()
	testReadFile()
}

func testArgs() {
	// 命令行参数，go run os/os.go a b c
	args := os.Args
	fmt.Println(args)
}

func testOpen() {
	// open file（流式读取）
	file, err := os.Open(`.gitignore`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open err, err: %+v\n", err)
		return
	}

	// scan file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// close file
	file.Close()
}

func testReadFile() {
	// 将文件内容一次性全部读取到内存
	data, err := os.ReadFile(`.gitignore`)
	if err != nil {
		log.Fatalf("ReadFile err, err: %+v\n", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}
