package main

import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}

	for i, t := range ts[:] { // 切片共享底层数组
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ") // 9
		}
	}

	fmt.Print(ts) //[{0} {9}]
}
