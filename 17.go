package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	// 将函数hello() 赋给h，并不是返回值
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	fmt.Printf("%T", h)
}
