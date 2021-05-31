package main

import "fmt"

type T struct {
	n int
}

func main() {
	m := make(map[int]T)
	//错误写法,map[key]struct 中 struct 是不可寻址的，所以无法直接赋值。
	//m[0].n = 1
	// 正确写法
	m[0] = T{1}
	fmt.Println(m[0].n)
}
