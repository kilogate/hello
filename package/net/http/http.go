package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	testGet()
}

func testGet() {
	// http get
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatalf("Get err, err: %+v\n", err)
		return
	}
	defer resp.Body.Close()

	// resp status code
	statusCode := resp.StatusCode
	fmt.Println(statusCode)

	// resp body
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadlAll err, err: %+v\n", err)
		return
	}
	fmt.Println(string(respBodyBytes))
}
