package main

import "fmt"

func main() {

	score(65)

}

func score(score uint8) {

	switch {
	case score > 60 && score <= 80:
		fmt.Println("很不错，加油")
	case score > 80 && score <= 90:
		fmt.Println("优秀，棒极了")
	case score > 90 && score <= 100:
		fmt.Println("你他娘的真是个天才")
	case score < 60:
		fmt.Println("真是个小费费")
	default:
		fmt.Println("宠辱不惊")

	}
}
