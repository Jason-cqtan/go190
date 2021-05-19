package main

import "fmt"

type S struct {
	m string
}

func f() *S {
	return &S{"foo"}
}

func main() {
	p := f() // *f()
	fmt.Println(p.m)
}
