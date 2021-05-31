package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 字符串拼接
	str := ""
	str = "abc" + "123"
	str = fmt.Sprintf("abc%d", 123)
	str = strings.Join([]string{"123", "456"}, "")

	var buffer bytes.Buffer
	for i := 0; i < 1000; i++ {
		buffer.WriteString("a")
	}
	str = buffer.String()

	fmt.Println(str)
}
