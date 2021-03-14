package main

import (
	"fmt"
	"strings"
)

func main() {
	callback(1, add)
	loveIndex := strings.IndexFunc("我爱中国", findLoveZH)
	fmt.Println(loveIndex)
}

func add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func findLoveZH(r rune) bool {
	return r == '爱'
}
