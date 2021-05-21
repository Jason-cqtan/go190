package main

import "fmt"

type N int

func (n N) test() {
	fmt.Println(n)
}

// 当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象，与其
//绑定，以便在稍后执行时，能隐式第传入接收者参数。
func main() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n) //13

	f1() //11
	f2() //12
}
