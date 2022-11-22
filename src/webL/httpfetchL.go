package webL

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpfetchL() {
	res, err := http.Get("http://www.baidu.com") //获取并显示网页内容
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}
