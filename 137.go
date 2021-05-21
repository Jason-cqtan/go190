package main

func main() {
	var x interface{}

	var y interface{} = []int{3, 5}

	_ = x == x
	_ = x == y
	_ = y == y // panic 两个比较值的动态类型为同一个不可比较类型。
}
