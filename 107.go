package main

func main() {
	var s []int
	s = append(s, 1)

	// 不能对 nil 的 map 直接赋值，需要使用 make() 初始化。
	var m map[string]int
	//m = make(map[string]int)
	m["one"] = 1
}
