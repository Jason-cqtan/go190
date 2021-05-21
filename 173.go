package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	defer f.Close() // defer 语句应该放在 if() 语句后面，先判断 err，再 defer 关闭文件句柄。
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(f)
	println(string(b))
}
