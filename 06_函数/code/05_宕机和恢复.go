package main

import "fmt"

func test5(s string) { //触发panic
	switch s {
	case "v1":
	case "v2":
	default:
		panic(fmt.Sprint("invalid value %q", s))
	}
}

func test6(x int) { //defer的执行顺序
	fmt.Printf("f(%d)\n", x+0/x) //当f(0)被调用时，发生panic异常
	defer fmt.Printf("defer %d\n", x)
	test6(x - 1)
}

func test7(x int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v", p)
		}
	}()
	fmt.Printf("%d\n", x+0/x)
}

func main() {
	test5("v1")
	//test6(3) // f(3)-> f(2)-> f(1)-> defer 1-> defer 2-> defer 3
	test7(0)
}
