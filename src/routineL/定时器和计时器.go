package routineL

import (
	"fmt"
	"time"
)

// the way to go 14.5
func TicktoBoom() {
	tick := time.Tick(1e8)  //周期性的发送时间给通道 (按照指定频率处理请求)
	boom := time.After(5e8) // 只发送一次时间给通道
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return //关闭通道
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
