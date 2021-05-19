package main

import "fmt"

// 当且仅当动态值 和动态类型都为 nil 时，接口类型值才为 nil
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo(x)
}
