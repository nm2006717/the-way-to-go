package main

import "fmt"

func main() {

	fmt.Println(Season(11))

}

func Season(month uint8) string {

	switch month {
	case 12, 1, 2:
		return "冬天"
	case 3, 4, 5:
		return "春天"
	case 6, 7, 8:
		return "夏天"
	case 9, 10, 11:
		return "秋天"
	default:
		return "error"

	}
}
