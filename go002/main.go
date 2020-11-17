package main

import (
	"fmt"

	"./samples"
)

func main() {
	var i int
	i = 10
	fmt.Println(i)
	i = 20
	fmt.Println(i)
	fmt.Println("abc")

	var x [3]int
	x[0] = 1
	x[1] = 2
	fmt.Println(x)
	fmt.Println(samples.Exc1(100))
}
