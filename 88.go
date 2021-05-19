package main

import "fmt"

const i = 100

var j = 123

// 常量不同于变量的在运行期分配内存， 常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，所以常量无法寻址。
func main() {
	fmt.Println(&j, j)
	fmt.Println(&i, i)
}
