package main

import "fmt"

func main() {
	nums := []int{1, 2, 6, 7, 8, 5, 4, -7, 8, -2, 5, 3, 0, 99, 78, 56}

	insertSort(nums)

	fmt.Println(nums)
}

//插入排序
//func insertSort (nums []int) {
//
//	for i:=1;i<len(nums);i++{
//		insertVal:=nums[i]
//		insertIndex := i-1
//		for insertIndex >=0 && nums[insertIndex] > insertVal {
//			nums[insertIndex+1] = nums[insertIndex] //数据后移
//			insertIndex --
//		}
//		if insertIndex+1 != i {
//			nums[insertIndex+1] = insertVal
//		}
//	}
//}

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {

		insertValue := nums[i]
		insertIndex := i - 1

		for insertIndex >= 0 && nums[insertIndex] > insertValue {
			nums[insertIndex+1] = nums[insertIndex]
			insertIndex--
		}

		if insertIndex+1 != i {
			nums[insertIndex+1] = insertValue
		}
	}
}
