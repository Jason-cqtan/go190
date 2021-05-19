package main

import "fmt"

// channel 声明、赋值、读取
func main() {
	var ch chan int
	ch = make(chan int)
	go func() {
		ch <- 7
	}()
	fmt.Println(<-ch)
}
