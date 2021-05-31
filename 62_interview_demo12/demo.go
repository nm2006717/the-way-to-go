package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	q := make(chan int)
	mutex := new(sync.Mutex)
	go func(q chan<- int) {
		defer close(q)
		for i := 1; i <= 100; i++ {
			q <- i
		}

	}(q)
	for i := 0; i < 3; i++ {
		go func(wg *sync.Mutex, q <-chan int) {
			for {
				wg.Lock()
				if v, ok := <-q; ok {
					fmt.Println(v)
					mutex.Unlock()
				} else {
					mutex.Unlock()
					return
				}
			}
		}(mutex, q)
	}
	time.Sleep(time.Second) // 安全退出
}
