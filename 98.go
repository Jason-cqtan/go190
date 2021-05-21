package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	*Param
}

// map 需要初始化才能使用;
// 指针不支持索引。
func main() {
	s := new(Show)
	//s.Param["day"] = 2

	// 修复代码如下:
	p := make(Param)
	p["day"] = 2
	s.Param = &p
	tmp := *s.Param
	fmt.Println(tmp["day"])
}
