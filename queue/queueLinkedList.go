package queue

import (

)

var queueLinkedList = list.New()

func enqueue(val int) {
	queueLinkedList.PushBack(val)
}

func dequeue() (int, error) {
	if queueLinkedList.Len() == 0 {
		return -1, errors.New("queue is empty")
	}
	return queueLinkedList.Remove(queueLinkedList.Front()).(int), nil
}