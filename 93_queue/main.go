package main

import (
	"fmt"
	"my-queue/queue"
)

func main() {

	queue := mq.New(2)

	queue.EnQueue(&mq.Element{Value: 1, Next: &mq.Element{Value: 1, Next: nil}})
	queue.EnQueue(&mq.Element{Value: 2, Next: nil})
	queue.DeQueue()

	fmt.Println(queue.Header.Value)
}
