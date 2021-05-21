package main

import "fmt"


// 数值溢出
func main() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			count++
		}
		if m == -m {
			count++
		}
	}
	fmt.Println(count) // 4 当i==0或i==128会使n==-n m==-m
}
