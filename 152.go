package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover()) // 捕获1
	}()
	defer func() {
		defer fmt.Print(recover()) //捕获2
		panic(1)
	}()
	defer recover() // 无效
	panic(2)
}
