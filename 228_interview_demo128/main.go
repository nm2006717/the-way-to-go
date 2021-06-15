package main

import "fmt"

func main() {

	// 覆盖预定义的nil
	nil := 123
	fmt.Println(nil)

	//var _ map[string]int = nil
}
