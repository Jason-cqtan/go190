package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	fmt.Println(*p) // 访问全局指针变量p的值
}

func main() {
	p, err := foo() // p 重新定义
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
