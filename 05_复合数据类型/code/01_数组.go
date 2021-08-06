package main

import "fmt"

func test1() {
	var arr1 [4]uint32
	var arr2 = [2]int{1,2}
	arr3 := [...]int{3,4}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	const(
		K1 = iota*2
		K2
	)
	arr4 := [...]string{K1: "v1", K2: "v2"} //k:[0 1 2]
	fmt.Println(arr4)
	fmt.Println(K2,arr4[K2])
	for k,v := range arr4 {
		fmt.Println(k,v)
	}

	arr5 := [...]int{5:-1} 
	fmt.Println(arr5) //[0 0 0 0 0 -1]

	arr6 := [2]int{1,2}
	arr7 := [...]int{2, 1}
	fmt.Println(arr6 == arr7)

	arr8 := []byte("ab c") //[97 98 32 99]
	fmt.Println(arr8)
	for k,v := range arr8 {
		fmt.Println(k,v)
	}
}

func main() {
	test1()
}