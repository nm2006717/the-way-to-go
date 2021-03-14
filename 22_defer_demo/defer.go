package main

import "fmt"

func main() {
	//function1()

	//a()

	//f()

	doDBOperations()
}

func function1(){
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom\n")
}

func function2(){
	fmt.Printf("Function2:Deferred until the end of the calling function")
}

func a(){
	i:=0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func connectToDB(){
	fmt.Println("ok,connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok,disconnected form db")
}

func doDBOperations(){
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}