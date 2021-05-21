package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}


// 结构体类型 data 没有实现接口 printer。知识点:接口。
func main() {
	d1 := data{"one"}
	d1.print()
	var in printer = data{"two"}
	in.print()
}