package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	//for range 循环的时候会创建每个元素的副本，而不是每个元素的引用，所以 m[key] = &val 取的 都是变量val的地址
	for key, val := range slice {
		m[key] = &val
	}

	//0 -> 3
	//1 -> 3
	//2 -> 3
	//3 -> 3
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
