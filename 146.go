package main

import "fmt"

func main() {
	var k = 9
	for k = range []int{} {
		fmt.Printf("rang空切片%v\n" , k)
	}
	fmt.Println(k) //9

	for k = 0; k < 3; k++ {
		fmt.Printf("for小于3->%v\n" , k)
	}

	fmt.Println(k) //3

	for k = range (*[3]int)(nil) {
		fmt.Printf("rang ->%v\n" , k)
	}

	fmt.Println(k) //2
}
