package webL

//15.2
import (
	"fmt"
	"net/http"
	"strings"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):] //get everything after the /hello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):] //get everything after the /shouthello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

// 执行调用
func WebserverL() {
	// http.HandleFunc("/", HelloServer) //处理请求
	// err := http.ListenAndServe("localhost:8080", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err.Error())
	// }

	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.ListenAndServe("localhost:9999", nil) //网页服务器监听端口 9999
}

// 然后打开浏览器并输入 url 地址：`http://localhost:8080/world`，
// 浏览器就会出现文字：`Hello, world`，网页服务器会响应你在 `:8080/` 后边输入的内容。
