package main

import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}

	for i, t := range ts { // range数组 ts 的副本，所以 t.n = 3 的赋值操作不会影响原数组。

		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ") // 0
		}
	}

	fmt.Print(ts) //[{0} {9}]
}
