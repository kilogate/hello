package main

import (
	"fmt"
	"io"
	"net/http"
)

type HttpResult struct {
	URL        string
	Err        error
	StatusCode int
	RespBody   string
}

func main() {
	urlList := []string{
		"http://www.baidu.com",
		"http://www.alibaba.com",
		"http://www.tencent.com",
	}

	httpResultMap := ParallelBatchGet(urlList)

	for url, httpResult := range httpResultMap {
		fmt.Println(url)
		if httpResult.Err != nil {
			fmt.Printf("get failed, err: %+v\n", httpResult.Err)
		} else {
			fmt.Printf("get success, status code: %d, resp body len: %d\n", httpResult.StatusCode, len(httpResult.RespBody))
		}
		fmt.Println()
	}
}

// ParallelBatchGet 并行批量获取
func ParallelBatchGet(urlList []string) map[string]*HttpResult {
	if len(urlList) == 0 {
		return nil
	}

	resChan := make(chan *HttpResult)
	for _, url := range urlList {
		go doGet(url, resChan)
	}

	resMap := make(map[string]*HttpResult, len(urlList))
	for range urlList {
		httpResult := <-resChan
		resMap[httpResult.URL] = httpResult
	}
	return resMap
}

func doGet(url string, resChan chan<- *HttpResult) {
	resp, err := http.Get(url)
	if err != nil {
		httpResult := &HttpResult{
			URL: url,
			Err: err,
		}
		resChan <- httpResult
		return
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		httpResult := &HttpResult{
			URL: url,
			Err: err,
		}
		resChan <- httpResult
		return
	}

	httpResult := &HttpResult{
		URL:        url,
		StatusCode: resp.StatusCode,
		RespBody:   string(respBodyBytes),
	}
	resChan <- httpResult
}
