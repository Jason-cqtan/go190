package main

import "fmt"

func main() {
	//var m map[string]int
	//m["a"] = 1
	//if v := m["a"]; v != nil {
	//	fmt.Println(v)
	//}

	m := make(map[string]int)
	m["a"] = 1
	if v, k := m["a"]; k != false {
		fmt.Println(v, k)
	}
}
