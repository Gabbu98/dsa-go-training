package queue

import (
	"container/list"
	"errors"
)

var queueLinkedList = list.New()

func enqueueLinkedList(val int) {
	queueLinkedList.PushBack(val)
}

func dequeueLinkedList() (int,error) {
	if queueLinkedList.Len() == 0 {
		return -1, errors.New("queue is empty")
	}
	return queueLinkedList.Remove(queueLinkedList.Front()).(int), nil
}