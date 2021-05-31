package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
	for i := 0; i < 4; i++ {
		v, ok := <-ch
		fmt.Println(v, ok)
	}

}
