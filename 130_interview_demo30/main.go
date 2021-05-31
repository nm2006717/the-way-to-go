package main

import "fmt"

type S struct {
	m string
}

func f() *S {
	return &S{"foo"}  //A
}

func main() {
	p := f()    //B
	fmt.Println(p.m) //print "foo"
}