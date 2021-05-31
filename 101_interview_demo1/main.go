package main

import "fmt"

func main() {
	var a MyInt = 1

	a.String()
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprint(m) //BUG:死循环
}



//for i := 1; i < 10; i++ {
//	go func(i int) {
//		fmt.Println(i)
//	}(i)
//}
//for {
//
//}