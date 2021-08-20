package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

//go run 01_flag_Value接口.go -period 5000ms
func main() {
	flag.Parse()
	fmt.Println(period)
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
