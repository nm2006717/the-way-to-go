package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 10000)
	for i := 0; i < 10000; i++ {
		abc <- i
	}
	defer close(abc)
	go func() {
		for {
			if a, ok := <-abc; ok {
				fmt.Println("a:", a)
			}
		}
	}()
	fmt.Println("close")
	time.Sleep(time.Millisecond * 10)
}
