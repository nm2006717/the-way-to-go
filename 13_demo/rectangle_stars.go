package main

import "fmt"

func main() {
	printRectangle()
}



func printRectangle() {
	for h := 0; h < 10; h++ {
		for w := 0; w < 20; w++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
