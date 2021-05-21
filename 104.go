package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}

	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	// for range 时会使用 a 的副本 a' 参与循 环，副本的 len 依旧是 5，因此 for range 只会循环 5 次，也就只获取 a 对应的底层数组的前 5 个元 素。
	fmt.Println(r) //[1,2,3,4,5]
}
