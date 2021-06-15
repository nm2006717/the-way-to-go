package main

import "fmt"

type N int

func (n N) test(){
	fmt.Println(n)
}

func main() {
	var n N =10
	fmt.Println(n)
	n++
	N.test(n)

	n++
	(*N).test(&n)
}