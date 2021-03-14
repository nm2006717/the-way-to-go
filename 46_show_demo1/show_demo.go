package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() *Student {
	var stu *Student
	fmt.Println(stu)
	return stu
}

func main() {

	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBBB")
	}
}
