package main

import "fmt"

func test(x byte) {
	fmt.Println(x)
}

// 别名类型无需转换，可直接转换
func main() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	test(c)
}
