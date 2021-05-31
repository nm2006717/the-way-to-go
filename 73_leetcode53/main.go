package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {

	max:=nums[0]

	for i:=1;i<len(nums);i++{
		if nums[i] + nums[i-1] >nums[i]{
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max

}

//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]`
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func main(){
	nums:=[]int{-2,1,-3,4}
	ret:= maxSubArray(nums)
	fmt.Println(ret)
}