package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	// RandomString：随机字符串
	randomString := lo.RandomString(10, lo.LettersCharset)
	fmt.Println(randomString)

	// Substring：字符串子串
	substring := lo.Substring("abcdef", 2, 3)
	fmt.Println(substring)

	// ChunkString：字符串分块
	chunkString := lo.ChunkString("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 5)
	fmt.Println(chunkString)

	// RuneLength：字符串字符数
	runeLength := lo.RuneLength("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fmt.Println(runeLength)
}
