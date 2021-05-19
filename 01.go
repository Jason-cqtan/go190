package main

import "fmt"

func main() {
	deferCall()
	// 后
	// 中
	// 前
	// panic
	//defer 的执行顺序是先进后出。出现panic语句的时候，会先按照 defer 的后进先出顺序执行，最后 才会执行panic。
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
