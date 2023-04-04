package main

import (
	"fmt"
	"time"

	"github.com/samber/lo"
)

func main() {
	toPtr := lo.ToPtr(time.Now())
	fmt.Println(toPtr)

	emptyableToPtr := lo.EmptyableToPtr(time.Now()) // 🤮
	fmt.Println(emptyableToPtr)

	fromPtr := lo.FromPtr(toPtr)
	fmt.Println(fromPtr)

	fromPtrOr := lo.FromPtrOr(toPtr, time.Now())
	fmt.Println(fromPtrOr)

	toSlicePtr := lo.ToSlicePtr([]string{"A", "B", "C"})
	fmt.Println(toSlicePtr)

	toAnySlice := lo.ToAnySlice([]any{"A", "B", "C", 1, 2, 3})
	fmt.Println(toAnySlice)

	fromAnySlice, ok := lo.FromAnySlice[any](toAnySlice)
	fmt.Println(fromAnySlice, ok)

	empty := lo.Empty[time.Time]()
	fmt.Println(empty)

	isEmpty := lo.IsEmpty(empty)
	fmt.Println(isEmpty)

	isNotEmpty := lo.IsNotEmpty(empty)
	fmt.Println(isNotEmpty)

	coalesce, ok := lo.Coalesce(0, 0, 7)
	fmt.Println(coalesce, ok)
}
