package main

import "fmt"

func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%v\n", a)
	ap(a)
	fmt.Printf("%v\n", a)
	app(a)
	fmt.Printf("%v\n", a)
}

// append 导致底层数组重新分配内存，a的slice底层数组不是外面一个，并没有改变外面这个
func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}
