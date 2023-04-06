package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	// Contains & ContainsBy：是否包含某个元素
	letters := []string{"A", "B", "C", "D", "C"}
	contains := lo.Contains(letters, "B")
	containsBy := lo.ContainsBy(letters, func(item string) bool {
		return item == "B"
	})
	fmt.Println(contains, containsBy)

	// Every & EveryBy：是否包含所有元素
	subLetters := []string{"A", "B", "C", "D", "C"}
	every := lo.Every(letters, subLetters)
	everyBy := lo.EveryBy(letters, func(item string) bool {
		return len(item) == 1
	})
	fmt.Println(every, everyBy)

	// Some & SomeBy：是否包含任意元素
	some := lo.Some(letters, subLetters)
	someBy := lo.SomeBy(letters, func(item string) bool {
		return len(item) == 1
	})
	fmt.Println(some, someBy)

	// None & NoneBy：是否不包含任意元素
	none := lo.None(letters, subLetters)
	noneBy := lo.NoneBy(letters, func(item string) bool {
		return len(item) == 1
	})
	fmt.Println(none, noneBy)

	// Intersect：交集
	otherLetters := []string{"A", "B", "C", "F"}
	intersect := lo.Intersect(letters, otherLetters)
	fmt.Println(intersect)

	// Difference：差集
	left, right := lo.Difference(letters, otherLetters)
	fmt.Println(left, right)

	// Union：并集
	union := lo.Union(letters, otherLetters)
	fmt.Println(union)

	// Without & WithoutEmpty：排除
	without := lo.Without(letters, "A", "C")
	withoutEmpty := lo.WithoutEmpty(letters)
	fmt.Println(without, withoutEmpty)
}
