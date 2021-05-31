package main

func main() {

}

func hasGroupsSizeX(deck []int) bool {
	maps := make(map[int]int)
	slices := make([]int, 0)
	for i := 0; i < len(deck); i++ {
		if _, ok := maps[deck[i]]; ok {
			maps[deck[i]] += 1
		} else {
			maps[deck[i]] = 1
		}

	}

	for _, value := range maps {
		if value <2 {
			return false
		}
		slices = append(slices, value)
	}

	if len(slices) >=2 {
		gcc:=gcc(slices)
		if gcc <2 {
			return  false
		}
		return true
	}

	return true

}

func gcc(nums []int) int {

	x := nums[0]
	y := nums[1]
	for y != 0 {
		x, y = y, x%y
	}
	if len(nums)>2{
		for i:=2;i<len(nums);i++{
			y = nums[i]
			for y !=0 {
				x,y = y,x%y
			}
		}
	}

	return  x
}
