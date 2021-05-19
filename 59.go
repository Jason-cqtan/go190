package main

import "fmt"

//func (i int) PrintInt() {
//	fmt.Println(i)
//}

type jason int

func (i jason) PrintInt() {
	fmt.Println(i)
}

func main() {
	//var i int = 1
	//i.PrintInt()
	var i jason = 1
	i.PrintInt()
}
