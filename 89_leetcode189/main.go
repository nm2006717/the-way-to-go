package main


func main(){
	nums:=[]int{1,2,3,4,5,6,7}
	k:=3

	rotate(nums,k)
}

func rotate(nums []int, k int)  {
	left:=nums[k:]
	right:=nums[:k]
	nums = append(left,right...)
}
