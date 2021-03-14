package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		defer func() {
			recover()
		}()
		for i := 0; i < 5; i++ {
			fmt.Println(5 / i)
			time.Sleep(1000)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1000)
	}

}
