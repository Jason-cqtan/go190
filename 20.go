package main

import "fmt"

type Person struct {
	name string
}

// 打印一个map中不存在的值时，返回元素类型的零值。
func main() {
	var m map[Person]int
	p := Person{"make"}
	fmt.Println(m[p])
}
