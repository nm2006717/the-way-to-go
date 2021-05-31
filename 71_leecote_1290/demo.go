package main

import "fmt"

func main() {
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	n := getDecimalValue(listNode)
	fmt.Println(n)

}

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var n int = 0
	list := head
	var sum int = 0
	for list != nil {
		n++
		list = list.Next
	}
	for head != nil {
		if head.Val == 1 {
			sum += 2<<n - 1
		}
		head = head.Next
		n--
	}
	return sum
}
