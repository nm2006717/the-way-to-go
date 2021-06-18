package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go func(done chan struct{}, ch chan int) {
		var resultValue int
		for {
			select {
			case resultValue = <-ch:
				fmt.Println(resultValue)
			case <-done:
				return
			}
		}

	}(done, ch)
	ch <- 1
	done <- struct{}{}
}
