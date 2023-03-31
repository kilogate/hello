package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	// ternary
	ternary := lo.Ternary(5 > 3, "if", "else")
	ternaryF := lo.TernaryF(5 > 3, func() string { return "if" }, func() string { return "else" })
	fmt.Println(ternary, ternaryF)

	// if else
	ifElse := lo.If(5 > 8, "if").ElseIf(6 > 8, "else if").Else("else")
	ifElseF := lo.IfF(5 > 8, func() string { return "if" }).ElseIfF(6 > 8, func() string { return "else if" }).ElseF(func() string { return "else" })
	fmt.Println(ifElse, ifElseF)

	// switch
	s := "F"
	switchCase := lo.Switch[string, string](s).
		Case("A", "is A").
		Case("B", "is B").
		CaseF("C", func() string { return "is C" }).
		Default("not match")
	fmt.Println(switchCase)
}
