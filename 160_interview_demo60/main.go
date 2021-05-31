package main

import "fmt"

var gvar int

func main() {
	var one int
	two := 2
	var three int
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}