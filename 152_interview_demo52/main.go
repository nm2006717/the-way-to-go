package main

import "fmt"

const i = 100
var j =123

// 常量无法寻址
func main() {
	fmt.Println(&j,i)
	fmt.Println(&i,i)
}
