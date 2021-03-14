package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Printf("the %dth fibonacci number is %d\n", i, f(i))
	}

}

func fibonacci() func(x int) int {
	var tmp int = 0
	return func(x int) int {
		back1, back2 := 0, 1
		for i := 0; i < x; i++ {
			back1, back2 = back2, (back2 + back1)
			tmp = back1
		}
		return tmp
	}
}
