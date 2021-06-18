package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c))

	x, ok := <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))

	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))

	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))

	//close(c)	// 此行将产生一个恐慌	close of closed channel
	//c <- 7	// 此行将产生一个恐慌	send on closed channel
}
