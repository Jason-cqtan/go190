package main

type S struct {
}

func f(x interface{}) {

}

func g(x *interface{}) {

}

// 函数参数为 interface{} 时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也用 interface{}，而不是使用 *interface{}。
func main() {
	s := S{}
	p := &s
	f(s)
	//g(s)
	f(p)
	//g(p)
}
