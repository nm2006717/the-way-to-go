package main

type X struct {
}

func (x *X) test() {
	println(x)
}

func main() {
	var a *X
	a.test()

	// X{}不可尋址
	//X{}.test()
}
