package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}

// hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时 用，所以是 5.

func main() {
	i := 5
	defer hello(i)
	i = i + 10
}
