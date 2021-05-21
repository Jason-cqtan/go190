package main

func main() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v) //2 2 2
		}()
		y[i] = &v
	}
	print(*y[0], *y[1], *y[2]) // 2 2 2
}
