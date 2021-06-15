package main

import "fmt"

type data struct {
	name string
}

type printer interface {
	print()
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

func main() {
	d1 := data{name: "one"}
	d1.print()

	// data结构体没有实现 printer接口

	//var in printer = data{name: "two"}
	// *data实现printer接口
	var in printer = &data{name: "two"}
	in.print()
}
