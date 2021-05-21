package main

import "fmt"

func main() {
	var k = 1
	var s = []int{1, 2}
	k, s[k] = 0, 3
	fmt.Println(s[0] + s[1]) // 4
}
