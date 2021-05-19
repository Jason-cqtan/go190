package main

import (
	"fmt"
)

// golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较。
func main() {
	var x string = nil
	if x == nil {
		x = "default"
	}

	fmt.Println(x)
}
