package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str1 = "asSASA ddd dsjkdsjs dk"
	var str2 = "asSASA ddd dsjkdsjsこん dk"
	fmt.Println(len(str1))
	fmt.Println(utf8.RuneCount([]byte(str1)))
	fmt.Println(len(str2))
	fmt.Println(utf8.RuneCount([]byte(str2)))

}
