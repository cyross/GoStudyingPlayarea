package main

import (
	"fmt"
)

const (
	S1 string = "a"
)

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()
	fmt.Println(f("hoge"))
	fmt.Println(f("fuga"))
	fmt.Println(f("fugafuga"))
	fmt.Println(S1)
}
