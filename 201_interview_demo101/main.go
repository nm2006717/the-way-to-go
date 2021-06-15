package main

import "fmt"

func main() {

	//	byte 是uint8的别名
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	test(c)

}

func test(x byte) {
	fmt.Println(x)
}
