package main

// 变量重复声明
func main() {
	one := 0
	one := 1

	//不能在单独的声明中重复声明一个变量，但在多变量声明的时候是可以的，但必须保证至少有一个变量
	//是新声明的。
	//two := 2
	//two, three := 2, 3
	//fmt.Println(two, three)
}
