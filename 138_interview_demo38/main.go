package main

import "fmt"

const (
	a = iota
	b = iota
)

const (
	name = "name"
	c    = iota
	d    = iota
)

//	知识点：iota 的用法。

//	iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。

//	iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
