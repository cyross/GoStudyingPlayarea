package main

import "fmt"

// 配列ポインタのデリファレンスを誤解なく理解する

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	p := &a

	fmt.Println((*p)[3]) // デリファレンスを明示
	fmt.Println(p[3])

	fmt.Println(len(*p)) // デリファレンスを明示
	fmt.Println(len(p))

	for i, v := range p {
		fmt.Printf("%d: %d\n", i, v)
	}
}
