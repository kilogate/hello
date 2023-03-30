package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	// IndexOf & LastIndexOf
	letters := []string{"A", "B", "C", "D", "C"}
	indexOf := lo.IndexOf(letters, "C")
	lastIndexOf := lo.LastIndexOf(letters, "C")
	fmt.Println(indexOf, lastIndexOf)

	// Find & FindIndexOf & FindLastIndexOf
	res, find := lo.Find(letters, func(item string) bool {
		return item == "D"
	})
	res, idx, find := lo.FindIndexOf(letters, func(item string) bool {
		return item == "D"
	})
	res, idx, find = lo.FindLastIndexOf(letters, func(item string) bool {
		return item == "D"
	})
	fmt.Println(res, idx, find)

	// FindOrElse
	res = lo.FindOrElse(letters, "X", func(item string) bool {
		return item == "Z"
	})

	// FindKey
	letterMap := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
		"d": "D",
	}
	key, ok := lo.FindKey(letterMap, "C")
	fmt.Println(key, ok)

	// FindUniques & FindDuplicates
	uniques := lo.FindUniques(letters)
	duplicates := lo.FindDuplicates(letters)
	fmt.Println(uniques, duplicates)

	// Min & MinBy & Max & MaxBy
	min := lo.Min(letters)
	minBy := lo.MinBy(letters, func(a string, b string) bool {
		return a < b
	})
	max := lo.Max(letters)
	maxBy := lo.MaxBy(letters, func(a string, b string) bool {
		return a > b
	})
	fmt.Println(min, minBy, max, maxBy)
}
