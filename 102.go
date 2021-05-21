package main

import "fmt"

// 全局参数未使用是允许的
var gvar int

func main() {
	//var one int
	//two := 2
	//var three int
	//three = 3
	// 函数参数未使用是允许的
	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}
