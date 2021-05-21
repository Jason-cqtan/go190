package main

// ch 为 nil，读写都会阻塞。
func main() {
	var ch chan int
	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}
}
