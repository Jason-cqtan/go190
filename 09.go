package main

import "fmt"

func main() {
	// 结构体比较
	// 结构体只能比较是否相等，不能比较大小
	// 相同类型的结构体才能进行比较，结构体是否相等不但与属性类型有关，还与属性顺序相关
	// 如果struct的所有成员都可以比较， 则该struct就可以通过==或!= 进行比较，比较时逐个项进行比较，如果每个项相等则两个结构体才相等，否则不相等
	// 能比较的：bool、数值型、字符、指针、数组
	// 不能比较：slice、map、函数
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "11"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
