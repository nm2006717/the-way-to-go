package main

import "fmt"

func main() {
	str := "你好，hello"
	a := []rune(str)
	a[0] = 'x'
	fmt.Println(string(a))
}
