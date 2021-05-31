package main

import "fmt"

func main() {
	nums := [14]int{22, 3, 5, 21, 20, 5, 4, 1, 0, 5, 6, 5, 5, 7}

	selectSort(&nums)

	fmt.Println(nums)
}

func selectSort(nums *[14]int) {
	n := len(nums)
	i := 0
	for i < n {
		j := i
		min := nums[i]
		minIndex := i
		for j < n {
			if min > nums[j] {
				min = nums[j]
				minIndex = j
			}
			j++
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
		i++
	}
}
