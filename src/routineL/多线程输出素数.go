package routineL

// the way to go 14.2
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

//协程 filter拷贝整数到输出通道，丢弃掉可以被 `prime` 整除的数字。
//然后每个 `prime` 又开启了一个新的协程，生成器和选择器并发请求。
// The prime sieve: Daisy-chain filter processes together.

// func main() {
// 	ch := make(chan int) // Create a new channel.
// 	go generate(ch)      // Start generate() as a goroutine.
// 	for {
// 		prime := <-ch
// 		fmt.Print(prime, " ")
// 		ch1 := make(chan int)
// 		go filter(ch, ch1, prime)
// 		ch = ch1
// 	}
// }
