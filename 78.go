package main

import "fmt"

type cool int

var i int = 1

func main() {
	j := cool(i)
	fmt.Println(j)
	fmt.Printf("%T", j)

}
