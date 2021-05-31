package main

import (
	"fmt"
	"os"
)

func main() {
	readFile()
}

func readFile() {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()

	file.Chdir()
}
