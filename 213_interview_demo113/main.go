package main

import "fmt"

func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	//	检查map是否含有某一元素，直接判断元素的值并不是一种合适的方式
	//if v := x["two"]; v == "" {
	//	fmt.Println("no entry")
	//}

	if _, ok := x["two"]; !ok {
		fmt.Println("no entry")
	}
}
