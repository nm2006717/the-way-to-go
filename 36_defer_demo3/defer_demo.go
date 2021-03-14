package main

import (
	"errors"
	"fmt"
)

func main() {
	err := MyFunc()
	fmt.Println(err)

	deferCall()
}

func MyFunc() error {
	err := errors.New("此处发生了错误")
	if err != nil {
		return err
	}

	// return 后面的defer没机会执行

	defer func() {
		if err != nil {
			fmt.Println("出错了，返回错误")
		}
	}()
	return nil
}

func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	panic("触发异常")
}
