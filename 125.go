package main

// 函数只能与 nil 比较。
func main() {
	var fn1 = func() {}
	var fn2 = func() {}

	if fn1 != fn2 {
		println("fn1 not equal fn2")
	}
}
