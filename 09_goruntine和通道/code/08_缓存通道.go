package main

import "fmt"

func query() string {
	response := make(chan string, 3)
	go func() { response <- "A" }()
	go func() { response <- "A" }()
	go func() { response <- "A" }()

	return <-response
}

func main() {
	res := query()

	fmt.Println(res)
}
