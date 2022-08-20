package main

import "fmt"

func main() {
	var value float64
	var unit string
	s := "36.8C"
	fmt.Sscanf(s, "%f%s", &value, &unit)

	fmt.Println(value)
	fmt.Println(unit)
}
