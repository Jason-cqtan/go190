package main

import "fmt"

// defer、返回值。
// increaseA()的返回参数是匿名，increaseB()的返回参数是具名
// 多个defer的执行顺序为后进先出
// 所有函数在执行RET返回指令之前，都会检查是否存在defer语句，若存在则先逆序调用defer语句进行收尾工作再退出返回
// 匿名返回值是在return执行时被声明，具名返回值则是函数声明的同时被声明，因此在defer语句中只能访问具名返回值，而不能直接访问匿名返回值
// return 其实应该包含前后两个步骤：第一步是给返回值赋值（若为具名返回值则直接赋值，若为匿名返回值则先声明再赋值）；第二步是调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句，最后RET携带返回值退出函数；
//defer、return、 返回值三者的执行顺序应该是：return最先给返回值赋值；接着defer开始执行一些收尾工作，最后RET指令携带返回值退出函数

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}
