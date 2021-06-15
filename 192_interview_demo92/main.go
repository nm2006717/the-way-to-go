package main

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func main() {

	// fmt.Sprintf底层调用String()方法
	c := &ConfigOne{}
	c.String()
}
