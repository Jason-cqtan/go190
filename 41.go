package main

import "fmt"

// 代码中 A B 两处应该怎么修改才能顺利编译?
func main() {
	//var m map[string]int // A
	//m["a"] = 1
	//if v := m["a"]; v != nil { // B
	//	fmt.Println(v)
	//}

	m := make(map[string]int)
	m["a"] = 1
	if v, k := m["a"]; k != false {
		fmt.Println(v, k)
	}
}
