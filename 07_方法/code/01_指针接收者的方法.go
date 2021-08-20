package main

import "fmt"

//测试1--------
type point struct {
	x float64
	y float64
}

func (p *point) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func test1() {
	p1 := &point{
		x: 1.2,
		y: 2.1,
	}
	p2 := point{
		1.2,
		2.1,
	}
	p1.scaleBy(2.0)
	(&p2).scaleBy(2.0)
	//p2.scaleBy(2.0) //这里会进行&p2的隐式转换
	fmt.Println(p1)
	fmt.Println(p2)
	//point{1.2, 2.1}.scaleBy(2.0) //编译错误-不能获取字面量point{1.2, 2.1}的地址
}

//2测试2--------
type pStruct struct{}

func (pStruct) scaleBy() {}

type pInt int

func (pInt) scaleBy() {}

type pPoint *int

//func(pPoint) scaleBy(){} //编译错误-不允许本身是指针的类型进行方法声明

var pInt2 int

//func (pInt2) scaleBy() {} //编译错误-只有命名类型与指向他们的指针是可以出现在接收者申明处的类型

//3 三种形式 --------

//形式1：接收器的实参和形参都是类型T或者都是类型 *T
type point2 struct {
	x float64
	y float64
}

func (p point2) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}
func test2() {
	p1 := point2{
		x: 2.2,
		y: 3.3,
	}
	p1.scaleBy(2.0)
}

// 形式2：接收器实参是类型T，但接收器形参是类型 *T ，这种情况下编译器会隐式地为我们取变量的地址
type point3 struct {
	x float64
	y float64
}

func (p *point3) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}
func test3() {
	p1 := point3{
		x: 2.2,
		y: 3.3,
	}
	p1.scaleBy(2.0)
}

// 形式3：接收器实参是类型 *T ，形参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量
type point4 struct {
	x float64
	y float64
}

func (p point4) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}
func test4() {
	p1 := &point4{
		x: 2.2,
		y: 3.3,
	}
	p1.scaleBy(2.0)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
