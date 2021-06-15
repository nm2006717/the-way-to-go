package main

func main() {
	var s []int
	s = append(s, 1)

	//	不能对nil的map直接赋值，需要使用make()初始化
	var m map[string]int
	m["one"] = 1
}
