package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	// f := w.(*os.File)
	// c := w.(*bytes.Buffer)
	// fmt.Println(f) //&{0xc000076280}
	// fmt.Println(c) // panic

	//Comma-ok断言
	f, okf := w.(*os.File)
	c, okc := w.(*bytes.Buffer)
	fmt.Println(okf, f) //true,&{0xc000076280}
	fmt.Println(okc, c) //false,nil(*bytes.Buffer的零值)

	var n1 interface{} = 12
	n2, okn := n1.(int)
	fmt.Println(n2, okn)
}
