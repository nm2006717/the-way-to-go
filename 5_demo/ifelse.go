package main

import "fmt"

func main() {

	var score = 100

	if score >= 60  {
		fmt.Println("pass")
	} else if score > 60 && score <= 80 {
		fmt.Println("good job")
	} else if score > 80 && score <= 100 {
		fmt.Println("god job")
	}

}
