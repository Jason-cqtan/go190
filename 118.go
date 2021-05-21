package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x) // 1
	{
		fmt.Println(x) // 1
		i, x := 2, 2
		fmt.Println(i, x) // 2 2
	}
	fmt.Println(x) // 1
}
