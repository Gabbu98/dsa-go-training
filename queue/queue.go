package queue

import "errors"

var queue []int

func enqueue(val int) {
	queue = append(queue,val)
}

func dequeue() (int, error) {
	if len(queue) == 0 {
		return -1, errors.New("queue is empty")
	}
	tmp := queue[0]
	queue = queue[1:len(queue)]
	return tmp, nil
}