package main

import (
	"fmt"
)

func main() {


	nums :=[]int{1,2,3,4,4,3,2,1}

	isBool:= hasGroupsSizeX(nums)
	fmt.Println(isBool)

}


func hasGroupsSizeX(deck []int) bool {
	maps:=make(map[int]int)
	for i:=0;i<len(deck);i++{
		if _,ok:=maps[deck[i]];ok{
			maps[deck[i]]+=1
		}else{
			maps[deck[i]] = 1
		}

	}

	min := maps[deck[0]]

	for _,value :=range maps {
		if min > value{
			min = value
		}
	}


	if min <= 1 {
		return false
	}

	for _,value:=range maps{
		if value % min != 0 {
			return false
		}
	}

	return true
}