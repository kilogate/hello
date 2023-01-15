package main

import "fmt"

func main() {
	testFallThrough()
}

func testFallThrough() {
	strSlice := []string{"a", "b", "c"}

	for _, str := range strSlice {
		fmt.Printf("Round %s\n", str)

		switch str {
		case "a":
			fmt.Println("a")
			fallthrough
		case "b":
			fmt.Println("b")
		case "c":
			fmt.Println("c")
		default:
			fmt.Println("default")
		}
	}
}
