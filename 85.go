package main

import "fmt"

func main() {
	// 整型切片初始化
	s := []int{1, 2, 3, 4, 5}
	s = make([]int, 0)
	s = make([]int, 5, 10)
	fmt.Println(s)
}
