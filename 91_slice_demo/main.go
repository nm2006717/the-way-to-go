package main

import "fmt"

func main() {
	nums := make([]int, 2)
	fmt.Println(nums)
	changeSlice(nums)
	fmt.Println(nums)

	fmt.Printf("%p\n", nums)

	str := "helloworld"
	fmt.Printf("%p\n", &str)
	changeString(str)
}

func changeSlice(nums []int) {
	nums[1] = 1
	nums = append(nums, 2)
	fmt.Println(nums)

	fmt.Printf("%p\n", nums)
}

func changeString(s string) {
	fmt.Printf("%p\n", &s)
}
