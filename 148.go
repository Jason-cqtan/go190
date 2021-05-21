package main

import "fmt"

func main() {
	var x int8 = -128
	var y = x / -1
	fmt.Println(y) // -128 溢出
}
