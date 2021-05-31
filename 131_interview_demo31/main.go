package main
import (
	"fmt"
)
func main() {
	//var x string = nil	//	string零值为""
	var x string
	if &x == nil {
		x = "default"
	}
	fmt.Println(x)
}