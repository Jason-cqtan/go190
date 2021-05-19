package main

import (
	"fmt"
	"runtime"
)

// select 会随机选择一个可用通道做收发操作，所以可能触发异常，也可能不会。
func main() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}
}
