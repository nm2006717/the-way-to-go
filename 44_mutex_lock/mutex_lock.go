package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.RLock()
	defer ua.RUnlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := UserAges{ages: make(map[string]int)}
	var wg sync.WaitGroup
	wg.Add(20)
	for j := 0; j < 19; j++ {
		go func() {
			age := ua.Get("你妹")
			fmt.Println(age)
			wg.Done()
		}()
	}

	go func() {
		ua.Add("你妹", 18)
		wg.Done()
	}()
	wg.Wait()
}
