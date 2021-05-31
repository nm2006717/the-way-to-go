package main 

import ("fmt")

type Param map[string] interface{}

type Show struct {
	Param
}

func main1() {
	s:=new(Show)
	s.Param["RMB"] = 1000

	fmt.Println(s)
}