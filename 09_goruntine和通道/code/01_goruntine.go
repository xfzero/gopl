package main

import (
	"fmt"
	"time"
)

func f1() {
	fmt.Println("f1 start")
	time.Sleep(3)
	fmt.Println("f1 end")
}

func f2() {
	fmt.Println("f2 start")
	time.Sleep(4)
	fmt.Println("f2 end")
}

func main() {
	go f1()
	f2()
}
