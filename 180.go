package main

func alwaysFalse() bool {
	return false
}

// true Go 代码断行规则。
func main() {
	switch alwaysFalse()
	{
	case true:
		println(true)
	case false:
		println(false)
	}
}
