package main

import "fmt"

func main() {
	var fn1 = func() {}
	var fn2 = func() {}

	// 函数变量不能比较
	if fn1 != fn2 {
		println("fn1 not equal fn2")
	}
}
