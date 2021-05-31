package main

import "fmt"

func main() {
	isMatch := func(i int) bool {
		switch i {
		// 每个case完成后默认break,可以加上fallthrough,这样就会接着走下一个case语句，而不会跳出switch
		case 1:
		case 2:
			return true
		}
		return false
	}

	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}
