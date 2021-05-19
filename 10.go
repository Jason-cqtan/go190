package main

import "fmt"

func main() {
	// 通过指针变量p访问其成员变量name方式
	// p.name
	// (*p).name
	// & 取址运算符
	// * 指针解引用
	p := &struct {
		name string
	}{name: "jason"}

	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println((*p).name)
}
