package main

import "fmt"

type User struct {

}

type User1 User

type User2 =User

func (i User1) m1()  {
	fmt.Println("m1")
}


func (i User) m2()  {
	fmt.Println("m2")
}

func (i User2) m3()  {
	fmt.Println("m3")
}


func main() {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()

	var i User
	i.m3()

}