package main

import "fmt"

type TimesMatcher struct {
	base int
}

func NewTimesMatcher(base int) *TimesMatcher {
	return &TimesMatcher{base: base}
}

// Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会被释放(内存逃逸)，因此在 Go 语言中返回函数参数或临时变量是安全的。
func main() {
	p := NewTimesMatcher(3)
	fmt.Println(p)
}
