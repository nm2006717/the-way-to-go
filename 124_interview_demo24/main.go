package main

import "fmt"

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int  {
	return w.i + 10
}

func (w Work) ShowB() int  {
	return w.i + 20
}

func main() {
	c:=Work{3}
	var a A = c
	var b B = c
	//	 da、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，
	//	接口 B 也不包括方法 ShowA()，编译报错。看下编译错误：
	//fmt.Println(a.ShowB())
	//fmt.Println(b.ShowA())
}
