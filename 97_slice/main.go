package main

import "fmt"

func main() {
	b := make([]byte, 2048, 2048)
	fmt.Printf("b的地址:%p\n", b)
	handlerSlice2(b)
}

func handlerSlice(s []byte) (n int) {
	fmt.Printf("函数里的切片地址:%p\n", s)
	return copy(s, []byte{'a', 'b', 'c'})
}

func handlerSlice2(s []byte) (n int)  {
	s[0] ='a'
	fmt.Printf("函数里的切片地址:%p\n", s)
	fmt.Printf("s[0]%p\n\n", &s[1])
	return len(s)
}