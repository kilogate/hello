package main

import (
	"flag"
	"fmt"
	"strconv"
)

type Thousand int

func (t *Thousand) String() string {
	return strconv.Itoa(int(*t/1000)) + "K"
}

func (t *Thousand) Set(val string) error {
	newVal, err := strconv.Atoi(val)
	*t = Thousand(newVal)
	return err
}

func ThousandFlag(name string, value Thousand, usage string) *Thousand {
	f := value
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

var num = ThousandFlag("num", 0, "k")

// go run gopl/ch7/flag.go -num 8000
func main() {
	flag.Parse()
	fmt.Println(num)
}
