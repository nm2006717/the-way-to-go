package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {

	var poe People = new(Student)
	// var poe People = &Student{}
	think := "bitch"
	fmt.Println(poe.Speak(think))
}
