package main

import "fmt"

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	fmt.Println(n) // 10

	n++
	f1 := N.test
	f1(n) // 11

	n++
	f2 := (*N).test
	f2(&n) // 12
}
