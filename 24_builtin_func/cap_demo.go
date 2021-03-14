package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	fmt.Println(cap(ch))

}
