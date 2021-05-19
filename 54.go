package main

import "fmt"

var a bool = true

// defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能 执行后面的函数或方法。
func main() {
	defer func() {
		fmt.Println("1")
	}()

	if a == true {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()

}
