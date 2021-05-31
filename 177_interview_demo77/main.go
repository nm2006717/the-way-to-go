package main

import "fmt"


type N int

func (n N) value(){
	n++
	fmt.Printf("v:%p,%v\n",&n,n)
}

func (n *N) pointer(){
	*n++
	fmt.Printf("v:%p,%v\n",n,*n)
}


func main() {

	var a N = 25

	p := &a
	p1 := &p

	// Golang不知多级指针调用方法
	p1.value()
	p1.pointer()
}