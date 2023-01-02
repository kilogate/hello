package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	testCopy()
}

func testCopy() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatalf("Get err, err: %+v\n", err)
		return
	}

	// 避免申请缓冲区：直接复制没有中间商
	written, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("Copy err, err: %+v\n", err)
		return
	}

	fmt.Printf("\n\nwritten: %v\n", written)
}
