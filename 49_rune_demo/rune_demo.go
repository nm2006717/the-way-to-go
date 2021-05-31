package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var zhString = "我爱你，中国"
	var enString = "i love you,china"

	var runeString = []rune(zhString)

	fmt.Printf("the length of zhString is %v\n", utf8.RuneCountInString(zhString))
	fmt.Printf("the length of enString is %v\n", len(enString))
	fmt.Printf("the length of runeString is %v\n", len(runeString))
}
