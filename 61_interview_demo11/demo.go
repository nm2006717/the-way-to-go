package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	q := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup, q chan<- int) {
		defer wg.Done()
		defer close(q)
		for i := 1; i <= 100; i++ {
			q <- i
			time.Sleep(time.Microsecond) // 利用休眠实现对生产者的控制，但这种方案不可靠
		}

	}(wg, q)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, q <-chan int) {
			defer wg.Done()
			for {
				if v, ok := <-q; ok {
					fmt.Println(i, v)
				} else {
					return
				}
			}
		}(wg, q)
	}
	wg.Wait()
}
