package main

import "fmt"

type MyInt1 int   // 类型定义,创建了新类型MyInt1
type Myint2 = int // 别名

func main() {
	var i int = 0
	var i1 MyInt1 = MyInt1(i) // 需使用Myint1(i)强制转换
	var i2 Myint2 = i
	fmt.Println(i1, i2)
	fmt.Printf("%T\n", i1)
	fmt.Printf("%T\n", i2)
}
