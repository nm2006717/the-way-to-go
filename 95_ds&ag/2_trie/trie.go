package main

import "fmt"

type runeTireNode struct {
	child map[rune]*runeTireNode
	Value interface{}
}

type RuneTire struct {
	root *runeTireNode
}

func NewRuneTire() *RuneTire {
	return &RuneTire{root: &runeTireNode{child: make(map[rune]*runeTireNode)}}
}

func (r *RuneTire) Insert(key string, value interface{}) {
	node := r.root
	for _, c := range key {
		if n, ok := node.child[c]; !ok {
			newNode := &runeTireNode{child: make(map[rune]*runeTireNode)}
			node.child[c] = newNode
			node = newNode
		} else {
			node = n
		}
	}
	node.Value = value
}

func (r *RuneTire) Get(key string) interface{} {
	node := r.root
	for _, c := range key {
		if n, ok := node.child[c]; !ok {
			return nil
		} else {
			node = n
		}
	}
	return node.Value
}

func main() {
	r := NewRuneTire()
	r.Insert("河北", "河北")
	r.Insert("湖南", "湖南")
	r.Insert("湖北", "湖北省")
	r.Insert("a", "/n")
	fmt.Println(r.Get("a"))
}
