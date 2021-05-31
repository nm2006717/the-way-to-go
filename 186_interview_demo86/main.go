package main

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
	*T
}

func main() {
	s := S{}
	_ = s.foo
	s.foo()
	//_ = s.bar
}