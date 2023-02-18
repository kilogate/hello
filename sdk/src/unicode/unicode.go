package main

import (
	"fmt"
	"unicode"
)

func main() {
	isDigit := unicode.IsDigit('2')
	isLetter := unicode.IsLetter('a')
	upper := unicode.ToUpper('a')
	fmt.Println(isDigit, isLetter, upper)
}
