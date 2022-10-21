package main

import (
	"flag"
	"fmt"
	"goTes/pack"
	"log"
	"os"
	"runtime"
	"strings"
)

type exampleType struct {
	Hello string
}

func (e exampleType) World() string {
	return "world"
}

// 注：同一个目录下面不能有个多main
func main() {
	flag.Parse()
	var foo int = 123
	bar := 456 //两种声明语句等价 ;不带声明格式的只能在函数体中出现
	var ptr *int = &foo

	arr := []int{2, 3, 4} //切片
	fmt.Println(printsum(arr))
	fmt.Println(foo, bar, ptr)
	//接口
	// s := new(pack.Simple)
	// fmt.Println(pack.FI(s))

	//文件读入
	//pack.FileInput()
	//pack.FileInput()
	//这是版本1
}

func Setpcmd(c *pack.CMD) {
	c.Pre_cmd = "bash"
	c.Short_options = "-f"
}

func printsum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func switch_test() {
	var grade string
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
		fallthrough
	case 80:
		grade = "B"
	default:
		grade = "C"
	}
	fmt.Println(grade)

}

var where = func() { //闭包(匿名函数)进行调试
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

func CmdLineArgs() { //命令行参数
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	// os.Args[0]是程序本身的名字
	fmt.Println("Hi", who)

	args := os.Args[1:] //查看参数切片
	//fmt.Println(args)
	fmt.Println(args)

}
