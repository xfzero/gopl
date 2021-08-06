package main

import "fmt"

// 除法-整数相除舍弃小数(运算结果和参数运算的类型一致)
func test1() {
	fmt.Println(2.6/1.25)   				//2.08  
	fmt.Println(2.6/1.2)					//2.1666666666666665
	fmt.Println(float32(2.6)/float32(1.2))	//2.1666665
	fmt.Println(9/5)						//1
	
}

func main() {
	test1()
}
