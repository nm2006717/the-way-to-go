package main

import (
	"encoding/json"
	"fmt"
)

//	ch为nil，读写都会阻塞
type People struct {
	name string `json:"name"`
}

func main() {
	js := `{ 
	"name":"seekload"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}
