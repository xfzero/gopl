package main

import "fmt"

func test1() {
	s :="ab中c"

	fmt.Println(len(s)) //6
	fmt.Println(fmt.Sprintf("%c",s[1])) //b
	fmt.Println(s[2:5]) //中

	fmt.Println(s[:]) //ab中c
	fmt.Println(s[:2]) //ab
}

func test2() {
	s := `go aaa 
bbb
	ccc"ddd"
	`

	fmt.Println(s)
}

func test3() {
	for i,r := range "abcd中e" {
		fmt.Printf("%d\t%q\t%d\n",i, r, r)
	}

	fmt.Println(fmt.Sprintf("%c",20013)) //中
}

func test4() {
	s := "fgh了i"

	r := []rune(s)

	fmt.Println(r) //[102 103 104 20102 105]
}

func test5() {
	fmt.Println(string(65)) //A
}

func main() {
	test1()

	test2()

	test3()

	test4()

	test5()
}