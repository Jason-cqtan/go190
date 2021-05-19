package main

import "fmt"

type Person struct {
	age int
}

// 变量 person 是一个指针变量 。
//1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的 时候取出，即输出 28;
//2.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后 面的函数执行的时候取出仍是 28;
//
//3.闭包引用，输出 29;

func main() {
	person := &Person{28}

	// 1
	defer fmt.Println(person.age)

	// 2
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}
