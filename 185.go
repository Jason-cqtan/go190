package main

import "fmt"

// 多重赋值分为两个步骤，有先后顺序:
//  1.计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式;
//  2.赋值;
func main() {
	var a []int = nil
	a, a[0] = []int{1, 2}, 9

	//a = []int{1,2}
	//a[0] = 9

	fmt.Println(a)
}
