package main

func main()  {
	s:="aaa"

	largeGroupPositions(s)
}

func largeGroupPositions(s string) [][]int {
	start,end:=0,0
	ret  := make([][]int,0)
	for end < len(s)-1{
		end = end +1
		if s[end-1] != s[end] {
			if end - start >=2{
				group := make([]int,2)
				group[0] = start
				group[1] = end -1
				ret = append(ret[:],group)
			}
			start = end
		}
	}


	return ret
}
