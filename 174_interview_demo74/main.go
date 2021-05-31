package main

type X struct {

}

func (x *X) test()  {
	println(x)
}

func main() {
	var a *X
	a.test()

	// X{}是不可寻址的，不能直接调用方法
	//X{}.test()
}