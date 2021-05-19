package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 字符串拼接
	str := "abc" + "123"
	str = fmt.Sprintf("abc%d", 123)
	str = strings.Join([]string{"123", "456"}, "")

	var buffer bytes.Buffer
	buffer.WriteString("12356")
	str = buffer.String()

	fmt.Println(str)
}
