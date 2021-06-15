package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			fmt.Printf("i:%d\n", i)
		}(wg, i)
	}

	wg.Wait()
}
