package main

import "fmt"

// Go 语言中的字符串是只读的。编译失败
func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}
