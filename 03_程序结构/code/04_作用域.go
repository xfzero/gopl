package main

import (
	"fmt"
)

func f() uint32 {
    return 1
}

//局部变量覆盖了f()
func test1() {
    f := 2 
    //fmt.Println(f()) //编译报错
    fmt.Println(f)
}

//隐式词法块
func test2() {
    if x := f(); x == 0 {
        fmt.Println("x==0")
    } else if y := f(); y == x { //嵌套在第一个if语句中初始化部分声明的变量，在第二个语句中是可见的
        fmt.Println("y==x")
    } else {
        fmt.Println("else")
    }
    // fmt.Println(x,y) //编译错误，x和y在这里不可见
}

//如果想在后续使用的改进
func test3() {
    x := f()
    y := f()

    if x == 0 {
        fmt.Println("x==0")
    }
    fmt.Println(x,y)

    //或 将要使用的语句放在if块中
    if z := f(); z == 0 {
        fmt.Println("x==0")
        fmt.Println(x,y)
    }
}

func main() {
    test1()
    
    test2()    
}

