package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person:=&Person{28}

	defer fmt.Println(person.age)	//28

	defer func(p *Person) {			//28
		fmt.Println(p.age)
	}(person)


	defer func() {
		fmt.Println(person.age)		//29
	}()

	person = &Person{29}

}