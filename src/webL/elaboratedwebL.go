package webL

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var helloRequests = expvar.NewInt("hello-requests") //`expvar` 可以创建（Int，Float 和 String 类型）变量，并将它们发布为公共变量;通常它被用于服务器操作计数。

var webroot = flag.String("root", "D:\\", "web root directory")
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

// Simple counter server. POSTing to it will set the value.
type Counter struct {
	n int
}

type Chan chan int

func ElaboratedWebL() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(Logger))
	http.Handle("/go/hello", http.HandlerFunc(HelloServerCount))

	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)

	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot)))) // uses the OS filesystem 用文件系统响应请求 (路径转为文件路径)
	http.Handle("/flags", http.HandlerFunc(FlagServer))
	http.Handle("/args", http.HandlerFunc(ArgServer))

	http.Handle("/chan", ChanCreate())

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}

func Logger(w http.ResponseWriter, req *http.Request) { //   登录/的展示
	log.Print(req.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops 404 Not Found"))
}

func HelloServerCount(w http.ResponseWriter, req *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello, server!\n")
	//w.Write([]byte("hello, server!\n"))
	fmt.Fprintf(w, "counter hello = %d\n", helloRequests.Value())
}

func (ctr *Counter) String() string { return fmt.Sprintf("%d", ctr.n) }

// 计数器对象 `ctr` 有一个 `String()` 方法，所以它实现了 `expvar.Var` 接口。这使其可以被发布，尽管它是一个结构体。
// `ServeHTTP()` 函数使 `ctr` 成为处理器，因为它的签名正确实现了 `http.Handler` 接口。
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET": // increment n
		ctr.n++
	case "POST": // set n to posted value
		buf := new(bytes.Buffer)
		io.Copy(buf, req.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func FlagServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") //
	fmt.Fprint(w, "Flags:\n")                                   //`VisitAll()` 函数迭代所有的标签 (flag)，打印它们的名称、值和默认值
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

func ArgServer(w http.ResponseWriter, req *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}
func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) { //每当有新请求到达，通道的 `ServeHTTP()` 方法从通道获取下一个整数并显示。
	//由此可见，网页服务器可以从通道中获取要发送的响应，它可以由另一个函数产生（甚至是客户端）。
	timeout := make(chan bool)
	go func() {
		for {
			select {
			case msg := <-ch:
				io.WriteString(w, fmt.Sprintf("channel send #%d\n", msg))
			case <-timeout:
				io.WriteString(w, fmt.Sprint("End\n"))
				return
			}
			time.Sleep(1e9)
		}
	}()
	time.Sleep(8e9) //8秒后超时结束
	timeout <- true
}
