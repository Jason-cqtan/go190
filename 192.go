package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}
	s := a[1:2] // s => [1]
	s[0] = 11   // s => [11]
	s = append(s, 12)
	s = append(s, 13) // s=> [11,12,13]
	s[0] = 21
	fmt.Println(a) //[0, 11, 12]
	fmt.Println(s) //[21,12,13]
}
