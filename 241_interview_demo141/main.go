package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
		//	Add之后没有Done panic
		wg.Add(1)
	}()

	wg.Wait()
}
