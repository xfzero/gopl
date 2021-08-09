package main

import "fmt"

func test1() {
	type Employee struct {
		ID int
		Name, Addr string //相同类型的可以写在一行上
		Salary int
	}

	var dilbert Employee

	//通过点好方式访问成员
	dilbert.Salary += 5000
	fmt.Println(dilbert.Salary)

	//获取成员变量的地址，通过指针来访问
	addr := &dilbert.Addr
	*addr = "七宝"
	fmt.Println(dilbert.Addr) 

	//点好可以用在结构体指针上
	var emp *Employee = &dilbert
	emp.Addr += "万科" // 等价与(*emp).Addr += "万科"
	fmt.Println(emp.Addr)
}

func test2() {
	// var tree struct {
	// 	value int
	// 	left, right *tree //不可以定义tree,但可以定义指针，可以创建递归结构
	// }

	// fmt.Println(tree)
}

func test3() {
	type Point struct{X, Y int}
	p := Point{1, 2} //必须有顺序

	type Student struct{
		id uint32 //不可导出
		Name string
	}
	s := Student{Name:"jack"} //指定变量名可以不管顺序

	fmt.Println(p)
	fmt.Println(s)
}

func main() {
	test1()
	test2()
	test3()
}