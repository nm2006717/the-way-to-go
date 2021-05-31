package main

import "fmt"

var x = []int{2:2,3,4,0:1}

//	字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础之上加一，
//	所以输出[1 0 2 3]，而不是[1 3 2]。
func main() {
	fmt.Println(x)
}