package main

import "fmt"

func receiver(ch <-chan int) {
	fmt.Println("===start receiver===")

	for { // 無限ループ
		fmt.Printf("[S]receiver\n")
		v := <-ch // chのキューが空だとスリープ
		fmt.Printf("[E]receiver: %d\n", v)
	}
	fmt.Println("===end receiver===")
}

func main() {
	var c1 chan int

	fmt.Println("===start main===")

	c1 = make(chan int)

	fmt.Println("===created channel===")

	go receiver(c1)

	fmt.Println("===called goroutine===")

	for i := 0; i < 10; i = i + 1 {
		fmt.Printf("[S]main: %d\n", i)
		c1 <- i // このタイミングでreceiverが始まる
		fmt.Printf("[E]main: %d\n", i)
	}

	fmt.Println("===end main===")
}
