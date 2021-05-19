package main

import "fmt"

//nil只能赋值给指针、chan、func、interface、map、或slice、类型的变量。
func main() {
	var x interface{} = nil
	var y error = nil

	fmt.Println(x, y)
}
