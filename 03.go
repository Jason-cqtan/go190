package main

import "fmt"

func main() {
	// append 向slice追加元素，容量不为零0，slice前补零，否则创建
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	//[0 0 0 0 0 1 2 3]
	fmt.Println(s)
	//[1 2 3 4]
	sliceAppend()
}

func sliceAppend() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}
