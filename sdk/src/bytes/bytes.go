package main

import (
	"bytes"
	"fmt"
)

func main() {
	testBytes()
}

func testBytes() {
	fmt.Println(bytes.Contains([]byte{'a', 'b', 'c'}, []byte{'a', 'b'}))             // true
	fmt.Println(bytes.Count([]byte{'a', 'b', 'c', 'a', 'b', 'c', 'a'}, []byte{'a'})) // 3
	fmt.Println(bytes.Fields([]byte{'a', 'b', ' ', 'c'}))                            // [[97 98] [99]]
	fmt.Println(bytes.HasPrefix([]byte{'a', 'b', 'c'}, []byte{'a', 'b'}))            // true
	fmt.Println(bytes.HasSuffix([]byte{'a', 'b', 'c'}, []byte{'a', 'b'}))            // false
	fmt.Println(bytes.Index([]byte{'a', 'b', 'c'}, []byte{'a', 'b'}))                // 0
	fmt.Println(bytes.Join([][]byte{{'a'}, {'a'}, {'a'}}, []byte{','}))              // [97 44 97 44 97]

}
