package main

import "fmt"

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}
func main() {
	var data info
	// 不能使用短变量声明设置结构体字段值
	data.result, err := work()
	fmt.Printf("info: %+v\n", data)

	//var err error
	//data.result, err = work()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(data)
}
