package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("Go! Go!")
	fmt.Printf("CPU数： %d\n", runtime.NumCPU())
	fmt.Printf("goroutine数: %d\n", runtime.NumGoroutine())
	fmt.Printf("バージョン: %s\n", runtime.Version())
}
