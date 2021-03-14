package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := makeAddSuffix(".bmp")
	addJpeg := makeAddSuffix(".jpeg")

	fileBmp := addBmp("file")
	fileJpeg := addJpeg("file")
	fmt.Printf("The file name of fileBmp is %s\n", fileBmp)
	fmt.Printf("The file name of fileJpeg is %s\n", fileJpeg)
}

func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
