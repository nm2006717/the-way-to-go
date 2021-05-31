package main 

import "fmt"

type student struct {
	Name string 
}
func zhoujielun(v interface{}){
	switch msg:=v.(type){
	case *student:
		msg.Name = "awei"
	case student:
		msg.Name = "awei"
	}
}

func main() {
	student:=&student{
		Name:"zhoujielun",
	}
	zhoujielun(student)
	fmt.Println(student.Name)
}