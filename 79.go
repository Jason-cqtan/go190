package main

import "fmt"

//switch
// switch 单个case中，可以出现多个结果选项;
// 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case;
func main() {
	i := 0
	switch {
	case i == 0:
		fmt.Println(0)
		fallthrough
	case i == 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println(3)
	}
}
