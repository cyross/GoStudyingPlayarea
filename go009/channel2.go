package main

import "fmt"

// 単なる写経になっているので、検証コードを入れる

func send(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("チャネルから受け取った数:", v)
	}
}

func main() {
	c := make(chan int)
	go send(c)
	receive(c)

	fmt.Println("処理を終了します。")
}
