package main

import "fmt"

type T struct {
	n int
}

func getT() T {
	return T{}
}

func main() {
	// T{}无法寻址，不可直接赋值
	//getT().n = 1

	t := getT()
	p := &t.n

	*p = 1
	fmt.Println(t.n)
}
