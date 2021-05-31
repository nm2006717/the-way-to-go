package main

import "math"

func main() {

	A:=[]int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0}
	_=prefixesDivBy5(A)
}

func prefixesDivBy5(A []int) []bool {
	cur:=0
	ret:=make([]bool,0)
	tmp:=make([]int,0)
	for cur != len(A) {
		tmp = append(tmp,A[cur])
		sum :=0
		for i,j:=len(tmp)-1,0;i>=0;i--{
			if tmp[i] != 0{
				sum = sum + int(math.Pow(2,float64(j)))
			}
			j++
		}
		if sum % 5 == 0{
			ret = append(ret,true)
		}else{
			ret = append(ret,false)
		}
		cur++
	}

	return ret
}