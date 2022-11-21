// 生成器是指当被调用时返回一个序列中下一个值的函数 14.8
// 生成器每次返回的是序列中下一个值而非整个序列；
// 这种特性也称之为惰性求值：只在你需要时进行求值，同时保留相关变量资源（内存和 CPU
package routineL

var Resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-Resume
}

// func main() {
// 	resume = integers()
// 	fmt.Println(generateInteger()) //=> 0
// 	fmt.Println(generateInteger()) //=> 1
// 	fmt.Println(generateInteger()) //=> 2
// }
