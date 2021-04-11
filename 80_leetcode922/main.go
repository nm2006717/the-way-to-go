package main

import "math"

func main(){

	nums:=[]int{3,1,4,2}

	sortArrayByParityII(nums)
}

func sortArrayByParityII(nums []int) []int {
	i,j:=0,1
	n:=len(nums)
	math.Pow10(10)
	for i<n&&j<n {
		if nums[i] % 2 !=0 {
			if nums[j] %2 == 0 {
				nums[i],nums[j] = nums[j],nums[i]
			}else{
				j =j+2
			}
		}else{
			i = i+2
		}
	}

	return nums
}