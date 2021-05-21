package main

import (
	"fmt"
	"testing"
)

func hello(num ...int) {
	num[0] = 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}
func main() {
	t := &testing.T{}
	Test13(t)
}
