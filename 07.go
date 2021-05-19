package main

import "fmt"

func main() {
	// append 第二个参数不能直接使用slice，需使用... 展开s2...,或者直接跟上元素
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
