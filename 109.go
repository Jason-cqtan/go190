package main

// nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。如果不指定变量的类型，编译器猜 不出变量的具体类型，导致编译错误。
func main() {
	var x = nil // var x interface{} = nil
	_ = x
}
