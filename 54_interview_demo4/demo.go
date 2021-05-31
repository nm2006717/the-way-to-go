package main

import (
	"fmt"
)

type People struct {
	Name string
}

func (p *People) String() string {
	// 回导致递归调用p.String()
	return fmt.Sprintf("print:%v", p)
}

func main() {
	p := &People{}
	p.String()
}
