package main

import "fmt"

// 当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil
func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println("not nil")
}
