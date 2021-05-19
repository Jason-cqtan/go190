package main

import "fmt"

// map 的输出是无序的。
//
//0 zero 1 one // 或者 1 one 0 zero
func main() {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
