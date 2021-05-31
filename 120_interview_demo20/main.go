package main

import "fmt"

func main() {
	p:=1
	incr(&p)
	fmt.Println(p)
}

func incr(p *int) int  {
	*p++
	return *p
}