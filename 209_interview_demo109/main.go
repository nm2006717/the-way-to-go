package main

import "fmt"

func main() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
			fallthrough
		case 2:
			return true
		}
		return false
	}

	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))

	match := func(i int) bool {
		switch i {
		case 1, 2:
			return true
		}
		return false
	}
	fmt.Println(match(1))
	fmt.Println(match(2))
}
