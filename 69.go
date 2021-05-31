package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n // 4+3 => 7
		recover()// 3
	}()

	var f func()

	defer f()//2 触发panic

	f = func() {
		r += 2
	}

	return n + 1//1 -> 4
}

func main() {
	fmt.Println(f(3))
}
