package main

// recover() 必须在 defer() 函数中直接调用才有效。
func main() {
	defer func() {
		recover()
	}()
	panic(1)
}
