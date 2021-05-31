package main

import "fmt"

type Person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	// new()与make()的区别
	// new(T) 和 make(T, args) 是Go语言内建函数，用来分配内存，但适用的类型不同。
	// new(T) 会为了 T 类型的新值分配已置零的内存空间，并返回地址(指针)，即类型为 *T 的 值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组 、 结构体 等。

	// make(T, args) 返回初始化之后的T类型的值，也不是指针 *T ，是经过初始化之后的T的引用。 make() 只适用于 slice 、 map 和 channel 。
	arr := new([5]int)
	var p *Person = new(Person)
	p.first = "first"
	p.last = "last"
	p.favFlavors = []string{"code", "game"}

	slice := make([]int, 5)
	m := make(map[int]*int)
	ch := make(chan int)

	fmt.Println(arr)
	fmt.Println(p)

	fmt.Println(slice)
	fmt.Println(m)
	fmt.Println(ch)
}
