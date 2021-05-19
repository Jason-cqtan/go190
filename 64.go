package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": {2, 3},
}

var n = map[string]*Math{
	"foo": &Math{2, 3},
}

func main() {
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp

	//m["foo"].x = 4
	fmt.Println(m["foo"].x)

	n["foo"].x = 4
	fmt.Println(n["foo"].x)
	fmt.Printf("%#v", n["foo"])
}
