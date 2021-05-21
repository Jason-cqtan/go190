package main

import "fmt"

func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	// 检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。最可靠的操作是使用访问 map 时返回的第二个值。
	if v := x["two"]; v == "" {
		fmt.Println("no entry")
	}

	//if _,ok := x["two"]; !ok {
	//	fmt.Println("no entry")
	//}
}
