package main

import (
	"fmt"
	"time"
)

func main() {
	response := make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++ {
			response <- i
			fmt.Println("丢入通道", i)
		}
	}()

	for {
		res := <-response
		fmt.Println(res)
		time.Sleep(1)
	}

}
