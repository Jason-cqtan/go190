package main

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {

	add(1, 2)
	add(1, 3, 7)
	add([]int{1, 3, 7}...)
}
