package main

import "fmt"

//截取操作有带 2 个或者 3 个参数，形如:[i:j] 和 [i:j:k]，假设截取对象的底层数组⻓度为 l。在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的⻓度，截取得到的切片⻓度和容量计算方法是 j- i、l-i。操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的⻓度 l，截取得到的切片⻓度 和容量计算方法是 j-i、k-i。

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	fmt.Println(len(a), cap(a)) //03
	fmt.Println(len(b), cap(b)) //23
	fmt.Println(len(c), cap(c)) //12
}
