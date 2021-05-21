package main

import "fmt"

var flag bool

func main() {
	if flag {
		fmt.Println("true")
	}

	if flag == false {
		fmt.Println("false")
	}

	if !flag {
		fmt.Println("!false")
	}
}
