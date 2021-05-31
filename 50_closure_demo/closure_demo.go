package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(12)
	for i := 0; i < 6; {
		time.Sleep(time.Second)
		go func() {
			fmt.Println("a:", i)
			wg.Done()
		}()
		i++
	}

	for i := 0; i < 6; i++ {
		go func(i int) {
			fmt.Println("b:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
