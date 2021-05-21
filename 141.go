package main

import "fmt"

type N int

// 当目标方法的接收者是指针类型时，那么被复制的就是指针。
func (n *N) test() {
	fmt.Println(*n)
}

func main() {
	var n N = 10
	p := &n
	n++
	f1 := n.test
	n++
	f2 := p.test
	n++
	fmt.Println(n) //13
	f1()           //13
	f2()           //13
}
