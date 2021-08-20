package main

import (
	"fmt"
	"time"
)

func main() {
	response := make(chan int, 1)

	go func() {
		for i := 0; i < 10; i++ {
			response <- i
			fmt.Println("丢入通道", i)
		}
	}()

	for {
		time.Sleep(1)
		res := <-response
		fmt.Println(res)
	}

}
