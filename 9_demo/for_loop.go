package main

import "fmt"

func main() {

	for i := 0; i < 15; i++ {
		fmt.Println("i = ", i)
	}
	a := 0
loop:
	a++
	fmt.Println("a = ", a)
	if a < 15 {
		goto loop
	}

}
