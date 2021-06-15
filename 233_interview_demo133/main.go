package main

import "fmt"

const (
	azero = iota
	aone  = iota
)

const (
	info  = "msg"
	bzero = iota
	bone  = iota
)

func main() {

	fmt.Println(azero, aone)
	fmt.Println(bzero, bone)
}
