package main

import (
	"fmt"
)

func main() {
	a := 0
	fn := func() int {
		a++
		return a
	}
	ch := make(chan int, 1)
	chh := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func(j int) {
			for {
				ch <- 1
				n := fn()
				if n > 100 {
					// 当第一个协程进入这段代码区域时，chh最终会被关闭
					// 剩下的两个协程因为chh被关闭，在写入按理来说应该会报panic: send on closed channel
					// 然而并未出现，原因在于其实协程在ch <- 1 的位置被阻塞了，为何?
					// ch容量为1，在第一个协程退出的时候并为将ch中的值取出，导致被阻塞
					chh <- 1
					close(chh)
					return
				}
				fmt.Println("go", j, n)
				<-ch
			}
		}(i)
	}
	for i := 0; i < 3; i++ {
		// 那这样是不是代表这chh中只有一个值，是的！但是为何主进程没阻塞呢？
		// 原因在于channle被关闭能读出0，所以这个循环在chh被关闭后马上就结束了
		<-chh
	}
}
