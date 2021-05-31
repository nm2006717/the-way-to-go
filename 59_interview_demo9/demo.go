package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0
	ch := make(chan int, 1)
	defer close(ch)
	for i := 0; i < 3; i++ {
		go func(j int) {
			for {
				ch <- 1
				a++
				if a > 10000 {
					break
				}
				fmt.Println("go", j, a)
				<-ch
			}
		}(i)
	}
	time.Sleep(time.Second * 8)
}
