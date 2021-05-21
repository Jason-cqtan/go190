package main

import "fmt"

type ConfigOne struct {
	Daemon string
}

// 如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用，最后抛错。
func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}
func main() {
	c := &ConfigOne{}
	c.String()
}
