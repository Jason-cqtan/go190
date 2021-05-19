package main

func main() {
	funcMui(1,2)
}

// 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括 号();如果只有一个返回值且命名也需要加上括号()。这里的第一个返回值有命名sum，第二个没有命名
func funcMui(x,y int) (sum int,error) {
	return x+y,nil
}
