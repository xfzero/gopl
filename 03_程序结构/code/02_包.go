package main

import (
	"fmt"
	//"D:\workspace\go\src\gopl\03_程序结构\code\ppackage"
	"./ppackage" //要关闭GO111MODULE(go env -w GO111MODULE=off),否则导包会有问题
)

func main() {
	name := ppackage.GetName()
	fmt.Println(name)
}