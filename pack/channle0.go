package pack

import (
	"fmt"
	"time"
)

func channle_test() {
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}

func sendDataCity(ch chan string) {
	ch <- "Washington"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}
func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func PowByChannel(s []int) {
	done := make(chan bool)
	// doSort is a lambda function, so a closure which knows the channel done:
	doPow := func(s []int) {
		for _, v := range s {
			v *= v
			fmt.Println(v)
		}
		done <- true
	}

	pivot := len(s) / 2
	go doPow(s[pivot:])
	go doPow(s[:pivot]) //多线程处理切片
	<-done
	<-done

}
