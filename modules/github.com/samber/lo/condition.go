package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	res := lo.Ternary(true, "IF", "ELSE")
	fmt.Println(res)
}
