package main

func main() {
	nums:=[]int{3,3,6,5,-2,2,5,1,-9,4}

	_=canThreePartsEqualSum(nums)
}

func canThreePartsEqualSum(arr []int) bool {
	sum:=0
	avg:=0
	n:=len(arr)
	for i:=0;i<n;i++{
		sum = sum+arr[i]
	}
	if sum % 3 != 0{
		return false
	}
	avg = sum / 3
	left,right:=1,n-2
	sumLeft := arr[0]
	sumRight:=arr[n-1]

	for left <= right {
		if sumLeft == avg {
			if sumRight == avg{
				return true
			}else{
				sumRight = sumRight + arr[right]
				right --
			}
		}else{
			sumLeft = sumLeft + arr[left]
			left++
			if sumRight!=avg{
				sumRight = sumRight + arr[right]
				right --
			}
		}
	}

	return false

}