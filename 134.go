package main

type T struct {
	n int
}

func getT() T {
	return T{}
}

func main() {

	// 直接返回的 T{} 无法寻址，不可直接赋值。

	getT().n = 1
	//t := getT()
	//t.n = 1
	//fmt.Println(t)
}
