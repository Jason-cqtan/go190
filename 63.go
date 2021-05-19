package main

import "fmt"

type Directon int

const (
	North Directon = iota
	East
	South
	West
)

func (d Directon) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(South)
}
