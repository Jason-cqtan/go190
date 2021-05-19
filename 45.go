package main

import "fmt"

//  接口的静态类型。

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

// 类型断言
func main() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Printf("%T\n", s)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
