package main

import "fmt"

func main() {
	t := NewTrie()
	t.Insert([]byte("hello"))
	t.Insert([]byte("world"))
	t.Insert([]byte("golang"))
	isExistHello := t.Search([]byte("world"))
	isStartWithHell := t.StartWith([]byte("holl"))
	fmt.Println(isExistHello)
	fmt.Println(isStartWithHell)
}

type Trie struct {
	IsEnd bool
	Next  [26]*Trie
	Value byte
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word []byte) {
	for _, value := range word {
		if t.Next[value-'a'] == nil {
			t.Next[value-'a'] = NewTrie()
			t.Value = value
		}
		t = t.Next[value-'a']
	}
	t.IsEnd = true

}

func (t *Trie) Search(word []byte) bool {
	for _, value := range word {
		if t.Next[value-'a'] == nil {
			return false
		}
		t = t.Next[value-'a']
	}
	return t.IsEnd
}

func (t *Trie) StartWith(word []byte) bool {
	for _, value := range word {
		if t.Next[value-'a'] == nil {
			return false
		}
		t = t.Next[value-'a']
	}
	return true
}
