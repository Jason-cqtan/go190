package main

import "fmt"

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

// 被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M。
func main() {
	f()
	fmt.Println("M")
}
