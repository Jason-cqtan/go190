package main

func main() {
	// 用字面量初始化数组、slice 和 map 时，最好是在每个元素后面加上逗号，即使 是声明在一行或者多行都不会出错。
	x := []int{
		1,
		2
	}
	_ = x
}