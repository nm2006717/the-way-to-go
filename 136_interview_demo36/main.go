package main

import "fmt"

//func (i int) PrintInt ()  {
//	fmt.Println(i)
//}
//
//func main() {
//	var i int = 1
//	i.PrintInt()
//
//}

type myInt int

func (mi myInt) PrintInt()  {
	fmt.Println(mi)
}

func main() {
	var i myInt = 1
	i.PrintInt()
}