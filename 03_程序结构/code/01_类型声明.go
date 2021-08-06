package main

import "fmt"

type pAge uint32
type aAge uint32

//类型声明
func test() {
	var pTony pAge = 18
	var pKity pAge = 18
	var aTony aAge = 2

	var num uint32 = 18

	//底层类型相同 不同类型声明的变量，不能进行运算
	fmt.Println(pTony, pKity, aTony, num)
	fmt.Println(pTony+pKity)
	//fmt.Println(pTony+aTony) //编译报错

	//可以和底层类型相同的未命名类型的值进行运算
	fmt.Println(pTony == 18)
	//fmt.Println(pTony == num) //编译报错

	
	fmt.Println(pTony.getAge())
	//fmt.Println(aTony.getAge()) //编译报错
}

//getAge的方法关联到pAge
func (this pAge) getAge() uint32 {
	return uint32(this)
}

func main() {
	test()
}