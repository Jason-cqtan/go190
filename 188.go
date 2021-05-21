package main

import "fmt"

const (
	azero = iota
	aone  = iota
)

// 在一个常量声明代码块中，如果 iota 没出 现在第一行，则常量的初始值就是非 0 值。
const (
	info  = "msg"
	bzero = iota
	bone  = iota
)

func main() {
	fmt.Println(azero, aone) // 0 1
	fmt.Println(bzero, bone) // 1 2
}
