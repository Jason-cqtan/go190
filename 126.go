package main

import "fmt"

type T struct {
	n int
}

func main() {
	m := make(map[int]T)
	// map[key]struct 中 struct 是不可寻址的，所以无法直接赋值
	m[0].n = 1

	//t := T{1}
	//m[0] = t

	fmt.Println(m[0].n)
}
