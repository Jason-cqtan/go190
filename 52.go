package main

import "fmt"

// A、B 两处应该填入什么代码，才能确保顺利打印 出结果?
type S struct {
	m string
}

func f() *S {
	return &S{"foo"} // A
}

func main() {
	p := f() // *f() // B
	fmt.Println(p.m)
}
