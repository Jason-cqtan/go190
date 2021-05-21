package main

import (
	"fmt"
	"runtime"
)

// for {} 独占 CPU 资源导致其他 Goroutine 饿死
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	for {
	}
}
