package main

var f = func(i int) {
	print("x")
}

func main() {
	f := func(i int) {
		print(i) // 10
		if i > 0 {
			f(i - 1) // x
		}
	}
	f(10)
}
