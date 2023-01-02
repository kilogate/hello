package main

import (
	"fmt"
	"io/ioutil"
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

	// resp status code
	statusCode := resp.StatusCode

	// println resp status code
	fmt.Println(statusCode)

	// resp body
	respBody := resp.Body

	// read resp body
	respBodyBytes, err := ioutil.ReadAll(respBody)

	// close resp body
	respBody.Close()

	if err != nil {
		log.Fatalf("ReadlAll err, err: %+v\n", err)
		return
	}

	// println resp body
	fmt.Println(string(respBodyBytes))
}
