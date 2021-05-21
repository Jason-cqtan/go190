package main

import "fmt"

func main() {
	nil := 123
	fmt.Println(nil)

	//  nil 是 int 类型值，不能赋值给 map 类型。
	var _ map[string]int = nil
}
