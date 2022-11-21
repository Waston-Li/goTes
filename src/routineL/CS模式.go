package routineL

import "fmt"

// //14.10
// 客户端-服务器应用正是 goroutines 和 channels 的亮点所在。
// 客户端 (Client) 可以是运行在任意设备上的任意程序，它会按需发送请求 (request) 至服务器。服务器 (Server) 接收到这个请求后开始相应的工作，然后再将响应 (response) 返回给客户端。典型情况下一般是多个客户端（即多个请求）对应一个（或少量）服务器。
//例如我们日常使用的浏览器客户端，其功能就是向服务器请求网页。而 Web 服务器则会向浏览器响应网页数据。
// 使用 Go 的服务器通常会在协程中执行向客户端的响应，故而会对每一个客户端请求启动一个协程。一个常用的操作方法是客户端请求自身中包含一个通道，而服务器则向这个通道发送响应。

// const MAXREQS = 50

// var sem = make(chan int, MAXREQS)  // 14.11 使用带缓冲的通道 限制同时处理的请求数
// 超过 `MAXREQS` 的请求将不会被同时处理，因为当信号通道表示缓冲区已满时 `run()` 函数会阻塞且不再处理其他请求，
// 直到某个请求从 `sem` 中被移除。`sem` 就像一个信号量，这一专业术语用于在程序中表示特定条件的标志变量。

type Request struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

type binOp func(a, b int) int //定义函数类型bin OP

func run(op binOp, req *Request) {
	//sem <- 1 // doesn't matter what we put in it
	req.replyc <- op(req.a, req.b)
	//<-sem // one empty place in the buffer: the next request can start
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service: // requests arrive here
			// start goroutine for request:
			go run(op, req) // don't wait for op
		case <-quit: //接受结束信号
			return
		}
	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func CSTest() {
	adder, quit := startServer(func(a, b int) int { return a + b })
	const N = 50
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req // adder is a channel of requests , 多个请求传入adder通道
	}
	// checks:
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}
	quit <- true
	fmt.Println("done")

}
