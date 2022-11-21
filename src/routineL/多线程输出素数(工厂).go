package routineL

// 函数 `sieve()`、`generate0()` 和 `filter()` 都是工厂；
// 它们创建通道并返回，而且使用了协程的 lambda 函数。
// Send the sequence 2, 3, 4, ... to returned channel
func generate0() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter0(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate0()
		for {
			prime := <-ch
			ch = filter0(ch, prime)
			out <- prime
		}
	}()
	return out
}

// func main() {
// 	primes := sieve()
// 	for {
// 		fmt.Println(<-primes)
// 	}
// }
