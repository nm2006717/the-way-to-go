package main

import "fmt"

type person struct {
	name string
}

func main() {
	var m map[person]int
	p :=person{"make"}
	fmt.Println(m[p])
}

