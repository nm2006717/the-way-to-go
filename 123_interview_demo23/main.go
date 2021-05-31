package main

import "fmt"

func main() {
	var m map[string]int     //A
	m = make(map[string]int) //	初始化map
	m["a"] = 1
	if v, ok := m["b"]; ok != false { //B
		fmt.Println(v)
	}
}
