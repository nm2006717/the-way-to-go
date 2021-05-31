package main

import "fmt"
import "strings"
import "strconv"
import "sort"

func main() {
	dates := []string{"20210520123", "20210420123", "线下20210520123", "20210520_123"}
	sort.Sort(NumsSort(dates))

	fmt.Println(dates)
}

type NumsSort []string

func (n NumsSort) Len() int {
	return len(n)
}

func (n NumsSort) Less(i, j int) bool {
	if strings.HasPrefix(n[i], "线下") && strings.HasPrefix(n[j], "线下") {
		intI, _ := strconv.Atoi(strings.Split(n[i], "线下")[1])
		intJ, _ := strconv.Atoi(strings.Split(n[j], "线下")[1])
		return intI < intJ
	}
	if strings.HasPrefix(n[i], "线下") && !strings.HasPrefix(n[j], "线下") {
		return true
	}
	if !strings.HasPrefix(n[i], "线下") && strings.HasPrefix(n[j], "线下") {
		return true
	}

	if !strings.HasPrefix(n[i], "线下") && !strings.HasPrefix(n[j], "线下") {
		intI, _ := strconv.Atoi(n[i])
		intJ, _ := strconv.Atoi(n[j])
		return intI < intJ
	}
	return false
}

func (n NumsSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
