package main

import "fmt"

// cap() 函数的适用类型 array、slice、channel

func main() {
	a := [2]int{1, 2}
	b := []int{1, 2}
	ch := make(chan int)
	fmt.Println(cap(a), cap(b), cap(ch))
}
