package main

type T struct{}

func (*T) foo() {
}
func (T) bar() {
}

type S struct {
	*T
}

func main() {
	s := S{} // main.S{T:(*main.T)(nil)}
	_ = s.foo
	s.foo()
	_ = s.bar// 空指针解析 panic
}
