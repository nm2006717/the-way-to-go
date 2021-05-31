package main

import "fmt"

func main() {

	nums := []int{0,4,0,3,2}
	k := 1
	max:= findMaxAverage(nums,k)
	fmt.Println(max)
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	if len(nums) == k{
		return float64(sum(nums))/float64(len(nums))
	}
	var current, max int
	slow, next := 0, k-1
	sumNums := make([]int, 0)
	sumNums = append(sumNums, nums[slow:next+1]...)
	max = sum(sumNums)
	slow ++
	next ++
	for next <= len(nums)-1 {
		current = current - nums[slow-1] + nums[next]
		if max < current {
			max = current
		}
		slow++
		next++
	}
	return float64(max)/float64(k)

}

func sum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}