package main

import (
	"encoding/json"
	"fmt"
)
//	name首字母小写，json包不能访问
type People struct {
	name string `json:"name"`
}

func main() {
	js:=`{
		"name":"seekload"
	}`
	var p People
	err:=json.Unmarshal([]byte(js),&p)
	if err!=nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Println(p)
}
