package main

import "fmt"

// 给一个 nil channel 发送数据，造成永远阻塞
// 从一个 nil channel 接收数据，造成永远阻塞
//  给一个已经关闭的 channel 发送数据，引起 panic
// 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
func main() {

	m := 1
	ch := make(chan int)
	go func() {
		ch <- m
	}()

	close(ch)
	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
}
