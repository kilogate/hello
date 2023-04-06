package main

import (
	"fmt"
	"time"

	"github.com/samber/lo"
)

func main() {
	// ToPtr：值 -> 指针（不支持参数为nil）
	toPtr := lo.ToPtr(time.Now())
	fmt.Println(toPtr)

	// EmptyableToPtr：值 -> 指针（支持参数为nil）
	emptyableToPtr := lo.EmptyableToPtr(time.Now()) // 🤮
	fmt.Println(emptyableToPtr)

	// FromPtr：指针 -> 值（参数为nil时返回零值）
	fromPtr := lo.FromPtr(toPtr)
	fmt.Println(fromPtr)

	// FromPtrOr：指针 -> 值（参数为nil时返回降级值）
	fromPtrOr := lo.FromPtrOr(toPtr, time.Now())
	fmt.Println(fromPtrOr)

	// ToSlicePtr：[]T -> []*T（不支持参数的元素为nil）
	toSlicePtr := lo.ToSlicePtr([]string{"A", "B", "C"})
	fmt.Println(toSlicePtr)

	toAnySlice := lo.ToAnySlice([]any{"A", "B", "C", 1, 2, 3})
	fmt.Println(toAnySlice)

	fromAnySlice, ok := lo.FromAnySlice[any](toAnySlice)
	fmt.Println(fromAnySlice, ok)

	empty := lo.Empty[time.Time]()
	fmt.Println(empty)

	// IsEmpty：是否零值
	isEmpty := lo.IsEmpty(empty)
	fmt.Println(isEmpty)

	// IsNotEmpty：是否零值
	isNotEmpty := lo.IsNotEmpty(empty)
	fmt.Println(isNotEmpty)

	coalesce, ok := lo.Coalesce(0, 0, 7)
	fmt.Println(coalesce, ok)
}
