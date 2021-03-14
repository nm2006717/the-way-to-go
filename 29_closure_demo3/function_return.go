package main

import (
	"fmt"
)

func main() {
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives:%v\n", p2(3))

	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	f := Adder3()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))

}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func Adder3() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
