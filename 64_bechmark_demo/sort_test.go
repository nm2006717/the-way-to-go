package main

import (
	"testing"
)

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(100000000)
	}
}

func Sort(num int) {
	// var arr []int
	for i := 0; i < num; i++ {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		bubbleSort(arr)
	}
}

//排序
func bubbleSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for k := 0; k < len(arr)-1-j; k++ {
			if arr[k] < arr[k+1] {
				temp := arr[k]
				arr[k] = arr[k+1]
				arr[k+1] = temp
			}
		}
	}
}
