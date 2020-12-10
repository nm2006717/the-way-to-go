package main

import "fmt"

func main() {
	//printGOne()
	printGTwo()
}

func printGOne() {
	for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("G")
		}
		fmt.Printf("\n")
	}
}

func printGTwo() {
	g := ""
	for i := 0; i < 25; i++ {
		g = g+"G"
		fmt.Println(g)
	}
}
