package main

import (
	"fmt"
	"sort"
)

func main() {

	nums:=[]int{-1,2,2,3,4,4}
	ret:= arraySign(nums)
	fmt.Println(ret)

}


//[-3,-2,-1,1,2,5,8]
//[-3] [5]
func arraySign(nums []int) int {
	sort.Ints(nums)
	left,right:=0,len(nums)-1
	for left <= right{
		pivot:=left+(right-left) >> 1
		if nums[pivot]>0{
			right = pivot-1
		}
		if nums[pivot]==0{
			return 0
		}
		if nums[pivot] <0{
			left = pivot+1
		}
	}
	if (left) %2==0{
		return 1
	}
	return -1
}
