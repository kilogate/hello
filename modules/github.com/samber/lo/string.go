package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	randomString := lo.RandomString(10, lo.LettersCharset)
	fmt.Println(randomString)

	substring := lo.Substring("abcdef", 2, 3)
	fmt.Println(substring)

	chunkString := lo.ChunkString("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 5)
	fmt.Println(chunkString)

	runeLength := lo.RuneLength("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fmt.Println(runeLength)
}
