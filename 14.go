package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

// iota 初始值为0，_标识不赋值
func main() {
	fmt.Println(x, y, z, k, p)
}
