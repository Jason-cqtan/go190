package main

import "fmt"

// 简短变量声明。使用简短变量声明有几个需要注意的地方:
//  只能用于函数内部;
//  短变量声明语句中至少要声明一个新的变量;
func main() {
	//  x 已声明，y 未声明
	var x int
	//x, _ = f()
	x, y := f()
	fmt.Println(x,y)
}

func f() (int,int) {
	return 0,0
}
