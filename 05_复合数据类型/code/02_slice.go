package main

import "fmt"

func test1() {
	month := [...]string{1:"Jan", 2:"Feb", 3:"Mar",4:"Apr", 5:"May", 6:"Jun"} //0:""
	Q1 := month[1:4]
	Sp := month[3:6]

	fmt.Println(Q1) //[Jan Feb Mar]
	fmt.Println(Sp) //[Mar Apr May]

	fmt.Println(Sp[:3]) //[Mar Apr May]
	fmt.Println(Sp[:4]) //[Mar Apr May Jun]  扩容
	fmt.Println(Sp)		//[Mar Apr May]
	//fmt.Println(Sp[:5]) //宕机
}

func test2() {
	s := "abc d"
	b := []byte(s)

	fmt.Println(s)
	fmt.Println(b)
}

func test3() {
	a := make([]int,3,5)
	fmt.Println(a)

	b := []rune("ab") //byte 等同于int8，即一个字节长度，常用来处理ascii字符
	fmt.Println(b)

	c := []byte("ab") //rune 等同于int32，即4个字节长度,常用来处理unicode或utf-8字符
	fmt.Println(c)
}

//重用底层数组
func nonempty(strings []string) []string {
	i := 0
	for _,s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

//模拟栈
func test4() {
	var stack []string

	stack = append(stack,"a")

	topStack := stack[len(stack)-1]

	fmt.Println(topStack)
}

func test5() {
	type st struct{
		name string
		age uint32
	}

	var s []st

	fmt.Println(s)
}

func main() {
	test1()
	test2()
	test3()

	data := []string{"a","b","","d"}
	data = nonempty(data) // 将data计算后重新赋值给data
	fmt.Println(data)

	test4()
	test5()
}