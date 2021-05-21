package main

import "fmt"

func main() {
	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k) //9

	for k = 0; k < 3; k++ {

	}

	fmt.Println(k) //3

	for k = range (*[3]int)(nil) {

	}

	fmt.Println(k) //2
}
