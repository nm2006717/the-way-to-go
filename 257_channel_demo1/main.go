package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func(ch chan<- int, x int) {
		ch <- x * x
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		done <- struct{}{}
	}(c)
	<-done
	fmt.Println("bye")
}
