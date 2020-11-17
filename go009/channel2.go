package main

import "fmt"

// 各文字のUnicode値を２増やして表示するプログラム
// goroutineとchannelを使用

func main() {
	str := "今日はいい天気"
	rstr := []rune(str)

	ch1 := make(chan rune)
	ch2 := make(chan rune)
	ch3 := make(chan rune)

	go func() {
		for {
			i := <-ch1
			ch2 <- i + 1
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- i + 1
		}
	}()

	ridx := 0
LOOP:
	for {
		// ch1 <- rstr[ridx] だと、 ridx==len(rstr)のときのガードができないため、
		// 処理を前倒ししてガードする方法しかなさそう
		// これだったらcloseを考えなくても良い
		var rr rune
		if ridx < len(rstr) {
			rr = rstr[ridx]
		} else {
			ch1 = nil
		}
		select {
		case ch1 <- rr:
			ridx++
		case r := <-ch3:
			fmt.Println(string(r))
		default:
			if ridx == len(rstr) {
				break LOOP
			}
		}
	}
}
