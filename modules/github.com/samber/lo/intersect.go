package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	// 包含某个元素：Contains & ContainsBy
	letters := []string{"A", "B", "C", "D", "C"}
	contains := lo.Contains(letters, "B")
	containsBy := lo.ContainsBy(letters, func(item string) bool {
		return item == "B"
	})
	fmt.Println(contains, containsBy)

	// 包含所有元素：Every
	subLetters := []string{"A", "B", "C", "D", "C"}
	every := lo.Every(letters, subLetters)
	everyBy := lo.EveryBy(letters, func(item string) bool {
		return len(item) == 1
	})
	fmt.Println(every, everyBy)

	// 包含任意元素：Some & SomeBy
	some := lo.Some(letters, subLetters)
	someBy := lo.SomeBy(letters, func(item string) bool {
		return len(item) == 1
	})
	fmt.Println(some, someBy)

	// 不包含任意元素：None & NoneBy
	none := lo.None(letters, subLetters)
	noneBy := lo.NoneBy(letters, func(item string) bool {
		return len(item) == 1
	})
	fmt.Println(none, noneBy)

	// 交集：Intersect
	otherLetters := []string{"A", "B", "C", "F"}
	intersect := lo.Intersect(letters, otherLetters)
	fmt.Println(intersect)

	// 差集：Difference
	left, right := lo.Difference(letters, otherLetters)
	fmt.Println(left, right)

	// 并集：Union
	union := lo.Union(letters, otherLetters)
	fmt.Println(union)

	// 排除：Without & WithoutEmpty
	without := lo.Without(letters, "A", "C")
	withoutEmpty := lo.WithoutEmpty(letters)
	fmt.Println(without, withoutEmpty)
}
