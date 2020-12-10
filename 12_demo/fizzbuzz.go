package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 != 0:
			fmt.Println("Fizz")
		case i%5 == 0 && i%3 != 0:
			fmt.Println("Buzz")
		case i%5 == 0 && i%3 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func printBizzBuzz() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
