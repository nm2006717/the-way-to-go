package main

import "fmt"

func main() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i is:", i) //0,1,2,
		i++
	}
	fmt.Println("A statement just after for loop")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
