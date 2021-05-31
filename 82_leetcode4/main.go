package main

func  main()  {

	num1:=[]int{4,5,6,8,9}
	num2:=[]int{}
	findMedianSortedArrays(num1,num2)
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1,n2:=len(nums1),len(nums2)
	i,j:=0,0
	var ret float64
	sortNums:=make([]int,0,2)
	if n1 == 0{
		if n2 % 2 == 0 {
			ret = float64(nums2[n2/2] + nums2[n2/2-1])/2.0
		}else{
			ret = float64(nums2[n2/2])
		}
		return ret
	}

	if n2 == 0 {
		if n1 % 2 == 0 {
			ret = float64(nums1[n1/2] + nums1[n1/2-1])/2.0
		}else{
			ret = float64(nums1[n1/2])
		}

		return ret
	}

	for {
		if nums1[i] > nums2[j] {
			sortNums = append(sortNums,nums2[j])
			if j <= n2-1{
				j++
			}
		}else{
			sortNums = append(sortNums,nums1[i])
			if i <= n1-1 {
				i++
			}
		}

		if j>= n2 {
			sortNums = append(sortNums,nums1[i:]...)
			break
		}
		if i>= n1 {
			sortNums = append(sortNums,nums2[j:]...)
			break
		}
	}

	if len(sortNums) % 2 == 0 {
		ret = float64(sortNums[len(sortNums)/2] + sortNums[len(sortNums)/2-1])/2.0
	}else{
		ret = float64(sortNums[len(sortNums)/2])
	}
	

	return ret

}
