package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a", a)
		}
	}()
	//管道在主协程被关闭了
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
