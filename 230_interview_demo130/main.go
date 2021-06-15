package main

import "fmt"

func main() {

	//	用于将内容从一个数组切片复制到另一个数组切片。如果加入的两个数组切片不一样大，
	//	就会按其中较小的那个数组切片的元素个数进行复制。
	dst := make([]int, 3)
	src := []int{1, 2, 3}

	copy(dst, src)
	fmt.Println(dst)
}
