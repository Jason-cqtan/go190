package main

import "fmt"

// := 只适用于函数内，必须使用显示初始化，不能提供数据类型
var (
	size := 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}