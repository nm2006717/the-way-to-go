package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 如果struct中存在飞导出的变量（以小写字母开头的变量名）不会被encode.
	bytes, _ := json.Marshal(Person{"你妹", 18})

	fmt.Println(string(bytes))

	p := new(Person)

	json.Unmarshal(bytes, p)

	fmt.Printf("p=%v\n", p)
	fmt.Printf("p=%+v\n", p)
	fmt.Printf("p=%#v\n", p)

}

type Person struct {
	name string
	Age  int
}
