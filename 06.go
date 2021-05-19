package main

import "fmt"

func main() {
	// new(T,arg) 返回指针，不能append
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)

}
