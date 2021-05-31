package main

import "sort"

func main() {
	//nums1:=[]int{1,2,3,0,0,0}
	//m :=3
	//nums2:=[]int{2,5,6}
	//n :=3
	//merge(nums1,m,nums2,n)

	nums3:=[]int{1,2,3,1,2,3}
	k :=2

	containsNearbyDuplicate(nums3,k)


}

// 双指针解法
func merge(nums1 []int, m int, nums2 []int, n int) {
    sorted := make([]int, 0, m+n)
    p1, p2 := 0, 0
    for {
        if p1 == m {
            sorted = append(sorted, nums2[p2:]...)
            break
        }
        if p2 == n {
            sorted = append(sorted, nums1[p1:]...)
            break
        }
        if nums1[p1] < nums2[p2] {
            sorted = append(sorted, nums1[p1])
            p1++
        } else {
            sorted = append(sorted, nums2[p2])
            p2++
        }
    }
    copy(nums1, sorted)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	tmp:=make([]int,0)
	for i:=0;i<len(nums)-k;i++{
		tmp = append(tmp[len(tmp):],nums[i:i+k+1]...)
		sort.Ints(tmp)
		for j:=0;j<len(tmp)-1;j++{
			if tmp[j] == tmp[j+1]{
				return true
			}
		}

	}
	return false
}