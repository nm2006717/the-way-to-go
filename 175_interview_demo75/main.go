package main

import "fmt"

func main() {
	x := map[string]string{"one":"a","two":"","three":"c"}

	// 错误，因为如果，x中不存在某个键值对是,v也会是""
	//if v := x["two"]; v == "" {
	//	fmt.Println("no entry")
	//}

	if _,ok:=x["two"];!ok{
		fmt.Println("no entry")
	}
}