package main

import "fmt"
type N int

func (n N) test(){
	fmt.Println(n)
}

func main()  {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	_ = y == y
}