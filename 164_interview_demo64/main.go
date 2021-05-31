package main

//	a 在 for range 过程中增加了两个元素，len 由 5 增加到 7，但 for range 时会使用 a 的副本 a' 参与循环，副本的 len 依旧是 5，因此 for range 只会循环 5 次，也就只获取 a 对应的底层数组的前 5 个元素。
func main() {
	var s []int
	s = append(s,1)

	//  map需要make后才能使用
	var m map[string]int
	m["one"] = 1
}
