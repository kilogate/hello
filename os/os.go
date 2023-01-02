package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testArgs()
	testFile()
}

func testArgs() {
	// 命令行参数，go run os/os.go a b c
	args := os.Args
	fmt.Println(args)
}

func testFile() {
	// open file
	file, err := os.Open(`.gitignore`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file fail, err: %+v\n", err)
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
