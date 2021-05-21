package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

// 1. Add() 方法的返回值依然是指针类型 *Slice，所以可以循环调用方法 Add();
// 2. defer 函数的参数(包括接收者)是在 defer 语句出现的位置做计算的，而不是在函数执行的时候
func main() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}
