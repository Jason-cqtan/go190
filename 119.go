package main

type foo struct {
	bar int
}

func main() {
	var f foo

	// := 操作符不能用于结构体字段赋值。
	f.bar, tmp := 1, 2
}
