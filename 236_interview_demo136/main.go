package main

import (
	"fmt"
	"sync"
)

//在协程中使用 wg.Add()；
//使用了 sync.WaitGroup 副本；
func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
