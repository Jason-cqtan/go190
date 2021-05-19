package main

import "fmt"

//a 的类型是int ，b 的类型是float ，两个不同类型的数值不能相加，编译报错。
func main() {
	a := 5
	b := 8.1
	fmt.Println(a + b)
}
