package main

import (
	"fmt"
	"strings"
)

//1
var f func(int) int

//2
func add1(v1 int) int {
	return v1 + 1
}

var f1 = add1

//3
func add2(r rune) rune { return r + 1 }
func test1() {
	fmt.Println(strings.Map(add2, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add2, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add2, "Admix"))
}

//4
func search(lists map[string]string, key string, help func(string, map[string]string) (result string, ok bool)) string {
	if val, ok := help(key, lists); ok {
		return val
	}
	return ""
}

func help1(v1 string, v2 map[string]string) (result string, ok bool) {
	val, ok := v2[v1]
	return val, ok
}

var fhelp1 = help1

func help2(v1 string, v2 map[string]string) (result string, ok bool) {
	val, ok := "", false
	for k, v := range v2 {
		if v == v1 {
			val = k
			ok = true
			break
		}
	}
	return val, ok
}

func test2() {
	student := map[string]string{
		"张三": "张大个",
		"李四": "小李飞刀",
	}

	if val, ok := search(student, "张三", fhelp1); ok { //报错
		fmt.Println(val)
	} else {
		fmt.Println("not found")
	}
}

func main() {
	test1()
	test2()

}
