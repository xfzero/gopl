package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x

			/*
				加上此判断，可以关闭naturals
				当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，
				它们会立即返回一个零值。
			*/
			if x >= 10 {
				break
			}
		}
		close(naturals)
	}()
	// Squarer
	go func() {
		for {
			// x := <-naturals
			x, ok := <-naturals
			if !ok {
				break
			}

			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
