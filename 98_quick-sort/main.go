package main

import (
	"fmt"
)

func main() {
	nums := []int{-9, 78, 0, 23, -567, 70}

	quickSort(0, len(nums)-1, nums)

	fmt.Println(nums)
}

func quickSort(left int, right int, nums []int) {
	l := left
	r := right
	pivot := nums[(left+right)/2]

	for ; l < r; {
		for ; nums[l] < pivot; {
			l++
		}
		for ; nums[r] > pivot; {
			r--
		}
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
		if nums[l] == pivot {
			r--
		}
		if nums[r] == pivot {
			l++
		}
	}

	if l == r {
		l++
		r--
	}

	if left < r {
		quickSort(left, r, nums)
	}
	if right > l {
		quickSort(l, right, nums)
	}

}
