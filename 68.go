package main

import (
	"fmt"
	"time"
)

// for range 使用短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重用， 而不是重新声明。
func main() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		//go func() {
		//	fmt.Println(i, v)
		//}()// 23 23 23

		// 传参
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)

		// 临时变量
		//i := i//这里的 := 会重新声明变量，而不是重用
		//v := v
		//go func() {
		//	fmt.Println(i, v)
		//}()
	}

	time.Sleep(time.Second * 3)
}
