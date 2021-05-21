package main

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}

func main() {
	println(DeferTest1(1)) //4
	println(DeferTest2(1)) //3
}
