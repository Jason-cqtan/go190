package main

import "fmt"

// 自增、自减不在是运算符，只能作为独立语句，而不是表达式;
// 不像其他语言，Go 语言中不支持 ++i 和 --i 操作;
func main() {
	data := []int{1, 2, 3}

	i := 0
	++i
	fmt.Println(data[i++])
}
