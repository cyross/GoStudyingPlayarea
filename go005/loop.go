package main

import "fmt"

func main() {
	fmt.Println("最初")
LOOP:
	for {
		for {
			for {
				fmt.Println("開始")
				break LOOP
			}
			fmt.Println("ここは通らない")
		}
		fmt.Println("ここは通らない")
	}
	fmt.Println("完了")
}
