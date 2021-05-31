package main

import "fmt"

func main() {
	nums:=[]int{2,3,99,45,11,10,11,2,5,5,4,3,1,0}
	bubbleSort(nums)

	fmt.Println(nums)
}

func bubbleSort(nums []int) {

	for i:=0;i<len(nums)-1;i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i] > nums[j] {
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
}


