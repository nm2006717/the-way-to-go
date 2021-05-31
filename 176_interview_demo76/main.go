package main

import "fmt"

type T struct {
	n int
}

func (t *T) Set(n int)  {
	t.n = n
}

func getT() T  {
	return T{}
}

func main() {
	//getT().Set(1)
	t := getT()
	t.Set(1)
	fmt.Println(t)
}