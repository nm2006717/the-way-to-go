package main

import (
	"sort"
)

func  main()  {
	nums:=[]int{5,7,7,8,8,10,11}
	target:=8

	searchRange(nums,target)
}


func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target + 1) - 1
	return []int{leftmost, rightmost}
}