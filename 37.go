package main

import "fmt"

//nil 切片和空切片。
//nil 切片和 nil 相等，一般用来表示一个不存在的切片;空切片和 nil 不相等，表示一个空的集合。
func main() {
	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
	fmt.Printf("%T\n", s2)
}
