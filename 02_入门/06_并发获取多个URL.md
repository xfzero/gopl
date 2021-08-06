1. goroutine
goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。
main函数本身也运行在一个goroutine中，而go function则表示创建一个新的goroutine，并在
这个新的goroutine中执行这个函数。

2. gopl.io/ch1/fetchall
```go
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

//总执行时间和执行时间最长的fetch有关
func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
            go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
            fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
            ch <- fmt.Sprint(err) // send to channel ch
            return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
            ch <- fmt.Sprintf("while reading %s: %v", url, err)
            return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
// ./fetchall https://golang.org http://gopl.io https://godoc.org https://www.baidu.com
```

当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调
用处，直到另一个goroutine往这个channel里写入、或者接收值，这样两个goroutine才会继续
执行channel操作之后的逻辑。在这个例子中，每一个fetch函数在执行时都会往channel里发
送一个值(ch <- expression)，主函数负责接收这些值(<-ch)。这个程序中我们用main函数来接
收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退
出。