package main

import "fmt"
import "sort"

func test1() {
	var a map[int]string

	b := map[int]string{
		1: "ab",
		2: "cd",
	}

	c := make(map[int]string)
	c[1] = "ab"
	c[2] = "cd"

	d := map[int]string{}
	d[1] = "ab"
	//a[1] = "ab" //空指针异常


	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	e := 12
	fmt.Println(&e) //0xc0000120f8
	fmt.Println(&c) //&map[1:ab 2:cd]
	// fmt.Println(&c[2]) //编译错误,其中的元素不是变量
}

func test2() {
	a := map[string]int{
		"k1": 1,
		"k2": 2,
	}

	fmt.Println(a["k1"])
	fmt.Println(a["k3"]) //0

	a["k2"] = 3
	fmt.Println(a)

	delete(a,"k2")
	fmt.Println(a)
}

func test3() {
	a := map[int]string{
		1: "v1",
		2: "v2",
		3: "v3",
		4: "v4",
		5: "v5",
	}

	//结果是随机的
	for k,v := range a {
		fmt.Println(k,v)
	}

	//实现顺序循环:将key生成一个对应的切片->对切片排序->对切片循环，更具切片循环的v取map中的元素
	b := make([]int, 0, len(a))
	for v := range a {
		b = append(b,v)
	}
	sort.Ints(b) //对b排序循环b
	for _,v := range b {
		fmt.Println(a[v])
	}
}

func main() {
	// test1()
	// test2()
	test3()
}
