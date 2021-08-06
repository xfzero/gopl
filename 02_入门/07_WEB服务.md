1. gopl.io/ch1/server1
```go
// Server1 is a minimal "echo" server.
package main
import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```
2. gopl.io/ch1/server2
```go
// Server2 is a minimal "echo" and counter server.
package main
import (
	"fmt"
	"log"
	"net/http"
	"sync"
)
var mu sync.Mutex
var count int
func main() {
	http.HandleFunc("/", handler) //其他请求都走这里 如：http://localhost:8000/count/123
	http.HandleFunc("/count", counter) //http://localhost:8000/count
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
```

这个服务器有两个请求处理函数，根据请求的url不同会调用不同的函数：对/count这个url的
请求会调用到counter这个函数，其它的url都会调用默认的处理函数。
如果你的请求pattern是以/结尾，那么所有以该url为前缀的url都会被这条规则匹配。
在这些代码的背后，服务器每一次接收请求处理时都会另起一个goroutine，这样服务器就可以同一时间处理多个请求。
然而在并发情况下，假如真的有两个请求同一时刻去更新count，那么这个值可能并不会被正确地增加；
这个程序可能会引发一个严重的bug：竞态条件（参见9.1）。
为了避免这个问题，我们必须保证每次修改变量的最多只能有一个goroutine，
这也就是代码里的mu.Lock()和mu.Unlock()调用将修改count的所有行为包在中间的目的。第九章中我们会进一步讲解共享变量。

3. gopl.io/ch1/server3
```go
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

```

在这些程序中，我们看到了很多不同的类型被输出到标准输出流中。比如前面的fetch程序，把HTTP的响应数据拷贝到了os.Stdout，
lissajous程序里我们输出的是一个文件。fetchall程序则完全忽略到了HTTP的响应Body，只是计算了一下响应Body的大小，
这个程序中把响应Body拷贝到了ioutil.Discard。
在本节的web服务器程序中则是用fmt.Fprintf直接写到了http.ResponseWriter中。
尽管三种具体的实现流程并不太一样，他们都实现一个共同的接口，即当它们被调用需要一
个标准流输出时都可以满足。这个接口叫作io.Writer

4. 结合gif
为了在这里简单说明接口能做什么，让我们简单地将这
里的web服务器和之前写的lissajous函数结合起来，这样GIF动画可以被写到HTTP的客户
端，而不是之前的标准输出流。只要在web服务器的代码里加入下面这几行。
```go
handler := func(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}
http.HandleFunc("/", handler)
```

