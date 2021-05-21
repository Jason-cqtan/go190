package main

import "fmt"

// 常量未使用是能编译通过
//的。
func main() {
	const x = 123
	const y = 1.23
	fmt.Println(x)
}
