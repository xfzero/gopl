package main
import (
	"fmt"
	"os"
)
//go run 01_命令行参数_输入.go 1 2 3
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:len(os.Args)] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[3])
}