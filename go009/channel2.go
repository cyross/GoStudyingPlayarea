package main

import "fmt"

func sendValue(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func receiveValue(c <-chan int) {
	for v := range c {
		fmt.Println("チャネルから受け取った数:", v)
	}
}

func main() {
	c := make(chan int)
	go sendValue(c)
	receiveValue(c)

	fmt.Println("処理を終了します。")
}
