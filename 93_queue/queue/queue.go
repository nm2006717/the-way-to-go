package mq

type Queue struct {
	Header *Element
	cap    int
}

type Element struct {
	Value interface{}
	Next  *Element
}

func New(cap int) *Queue{
	return  &Queue{nil,cap}
}

// 入队列
func (queue *Queue) EnQueue(ele *Element) *Element {
	if queue.canEnQueue() {
		header := queue.Header
		if header == nil {
			header = ele
			queue.Header = ele
			return ele
		}
		for header.Next != nil {
			header = header.Next
		}
		header.Next = ele
		return ele
	}
	return nil
}

// 出队列
func (queue *Queue) DeQueue() interface{} {
	if queue.isEmpty() {
		return nil
	}
	val := queue.Header.Value
	queue.Header = queue.Header.Next
	return val
}

//	判断队列是否为空
func (queue *Queue) isEmpty() bool {
	if queue.Header == nil {
		return true
	}
	return false
}

// 判断是否队列是否已满
func (queue *Queue) canEnQueue() bool {
	if queue.isEmpty() {
		return true
	}
	head := queue.Header
	cnt := queue.cap - 1
	for head.Next != nil {
		head = head.Next
		cnt--
	}
	if cnt > 0 {
		return true
	}
	return false
}
