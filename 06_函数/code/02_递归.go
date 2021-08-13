package main

import "fmt"

func test1(n int) int {
	if n == 1 {
		return 1
	}
	return test1(n-1) + n
}

func main() {
	fmt.Println(test1(3))
}
