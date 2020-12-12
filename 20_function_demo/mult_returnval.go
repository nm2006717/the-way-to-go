package main

import "fmt"

func main() {

	add, mult, sub := multReturnval(3, 5)
	fmt.Print("multRetunval.....")
	fmt.Printf("add=%d,mult=%d,sub=%d", add, mult, sub)
	fmt.Println()
	fmt.Print("multRetunval2....")
	add2, mult2, sub2 := multReturnval2(3, 5)
	fmt.Printf("add2=%d,mult2=%d,sub2=%d", add2, mult2, sub2)
}

func multReturnval(value1 int, value2 int) (int, int, int) {

	return value1 + value2, value1 * value2, value2 - value1

}

func multReturnval2(value1 int, value2 int) (add int, mult int, sub int) {
	add, mult, sub = value1+value2, value1*value2, value2-value1
	return
}
