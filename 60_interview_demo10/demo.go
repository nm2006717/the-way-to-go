package main

import (
	"fmt"
	"time"
)

func print(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}

}

func main() {
	ch0 := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer close(ch0)
	defer close(ch1)
	defer close(ch2)
	go print(ch0)
	go print(ch1)
	go print(ch2)
	for i := 1; i < 101; i++ {
		switch {
		case i%3 == 0:
			ch0 <- i
		case i%3 == 1:
			ch1 <- i
		case i%3 == 2:
			ch2 <- i
		}
		time.Sleep(time.Microsecond)
	}
}
