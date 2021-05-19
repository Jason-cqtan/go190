package main

import "fmt"

func GetValue() int {
	return 1
}

// 编译失败
// 只有接口类型才能使用类型选择 类型选择的语法形如:i.(type)，其中i是接口，type是固定关键字，需要注意的是，只有接口类型才可以 使用类型选择。
func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknow")

	}
}
