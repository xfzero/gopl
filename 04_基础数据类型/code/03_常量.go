package main

import "fmt"

//无类型常量
const pi = 3.14
const ret string = "ok"

const(
	NUM1 = iota
	NUM2
	NUM3 = iota
)
const NUM4 = iota

type Newtype uint32
const(
	NUM5 Newtype = iota
	NUM6
)

const(
	NUM7 = iota*2
	NUM8
)


const V1 = 10
var V2 = 10


func main() {
	fmt.Println(pi,ret)

	fmt.Println(NUM1,NUM2,NUM3,NUM4) //0 1 2 0

	fmt.Println(NUM5,NUM6)

	fmt.Println(NUM7,NUM8) //0 2

	var V3 uint32 = 10
	var V4 uint64 = 10
	fmt.Println(V1,V2,V3)
	fmt.Println(V1+V3)
	//fmt.Println(V2+V3) //编译报错，只有常量支出无类型，这里要转换数据类型
	fmt.Println(uint32(V2)+V3)
	fmt.Println(V1+V4)

	//变量声明了类型后，重新赋值时会隐式装换
	var a float64 = 3.14
	a = 4
	fmt.Printf("%T",a) //float64

}