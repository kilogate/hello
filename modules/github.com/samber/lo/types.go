package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	tuple2 := lo.Tuple2[string, string]{
		A: "AA",
		B: "BB",
	}
	fmt.Println(tuple2.Unpack())
}
