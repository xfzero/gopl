package main

import "fmt"

func test1() (string, error) {
	return "data", nil
}

//裸返回值
func test2() (result string, err error) {
	result, err = test1()
	return
}

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
}
