package main

import "fmt"

// iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。
//iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。
const (
	a = iota
	b = iota
)

const (
	name = "name"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a) //0
	fmt.Println(b) //1
	fmt.Println(c) //1
	fmt.Println(d) //2
}
