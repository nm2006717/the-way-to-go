package main

import "fmt"

func main() {
	stringElem := func(done <-chan interface{}) <-chan interface{} {
		stringValue := make(chan interface{})
		go func() {
			defer close(stringValue)
			select {
			case <-done:
				break
			case stringValue <- "hello":
			}
		}()
		return stringValue
	}

	done := make(chan interface{})
	defer close(done)
	for strVal := range stringElem(done) {
		fmt.Printf("%T\n", strVal)
		fmt.Printf("%v\n", strVal)
	}
}
