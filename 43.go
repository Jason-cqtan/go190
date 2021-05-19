package main

import "fmt"

// x 已声明 y 未声明
var x int

func main() {
	//x, _ := f()
	x, _ = f()
	x, y := f() // 当多值赋值时，:= 左边的变量无论声明与否都可以
	//x, y = f()
	fmt.Println(x, y)
}

func f() (int, int) {
	return 0, 1
}
