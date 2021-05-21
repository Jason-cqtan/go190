package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(recover()) // 捕获2
	}()
	defer func() {
		defer func() {
			fmt.Println(recover()) // 捕获1
		}()
		panic(1)
	}()
	defer recover() // 无效
	panic(2)
}
