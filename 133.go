package main

func main() {
	//  channel

	// 不能在单向通道上做逆向操作(例如:只发送通道用于接收); √
	// close() 可以用于只接收通道; ×
	// 单向通道可以转换为双向通道; ×
}
