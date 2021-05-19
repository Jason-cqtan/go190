package main

import "fmt"

// Go中的数组是值类型，可比较，另外一方面，数组的⻓度也是数组类型的组成部分，所以 a 和 b 是不 同的类型，是不能比较的，所以编译错误。
func main() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

}
