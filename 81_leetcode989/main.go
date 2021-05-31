package main

import "math"

func main(){
	num:=[]int{0}
	k :=0

	addToArrayForm(num,k)
}


func addToArrayForm(num []int, k int) []int {

	retNum:=make([]int,0)
	intNums:=0
	for i:=0;i<len(num);i++{
		intNums =intNums + num[i]* int(math.Pow(10.0,float64(len(num)-i-1)))
	}

	totalNum:=k+intNums

	for {
		retNum = append(retNum,totalNum % 10)
		totalNum = totalNum / 10
		if totalNum == 0{
			break
		}
	}
	reverse(retNum)

	return retNum
}


func reverse(A []int) {
	for i, n := 0, len(A); i < n/2; i++ {
		A[i], A[n-1-i] = A[n-1-i], A[i]
	}
}

