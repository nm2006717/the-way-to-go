package main

import (
	"fmt"
)

func test_recove() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recovers success")
			fmt.Println(err)
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}
func main() {
	test_recove()
	fmt.Println("after recover")
}
