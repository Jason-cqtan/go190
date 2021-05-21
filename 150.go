package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

// defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取 出。
func main() {
	f := F(5)
	defer func() {
		fmt.Println(f(), 3) //8
	}()

	defer fmt.Println(f(), 2) //6
	i := f()
	fmt.Println(i, 1) // 7
}
