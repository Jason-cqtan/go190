package main

import "fmt"

// A 声明的是 nil 切片;B 声明的是⻓度和容量都为 0 的空切片。第一种切片声明不会分配内存，优先选择。
func main() {
	// A
	var a []int
	// B
	b := []int{}
	fmt.Println(a, b)
}
