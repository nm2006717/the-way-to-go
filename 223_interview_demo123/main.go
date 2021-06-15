package main

type T struct {
}

type S struct {
	*T
}

func (*T) foo() {

}

func (T) bar() {

}

func main() {
	s := S{}
	_ = s.foo
	s.foo()

	//因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic。
	_ = s.bar
}
