package main

import "fmt"

func main(){
	nums:=[]int{18,66,24,29,57,54,84,16,29,71,22,33,98,34,20,65,21,6,11,74,7,93,40,64,97,56,3,72,63,67,72,86,42,29,34,46,97,34,45,82,22,80,94,45,46,96,10,34,98,55}
	index:= dominantIndex(nums)

	fmt.Println(index)
}


func dominantIndex(nums []int) int {
	max,second:=nums[0],0
	n:=len(nums)
	maxIndex:=0
	for i:=1;i<=n-1;i++{
		if nums[i] > max {
			second = max
			max = nums[i]
			maxIndex = i
		}
		if nums[i] < max {
			if second < nums[i] {
				second = nums[i]
			}
		}
	}
	if second == max {
		return 0
	}
	if second == 0 {
		return maxIndex
	}
	if max / second >= 2{
		return maxIndex
	}
	return -1
}