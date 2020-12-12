package main

import "fmt"

func main() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}

func greeting() {
	fmt.Println("In greeting: Hi!!!!!")
}
