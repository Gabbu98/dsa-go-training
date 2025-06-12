package queue

import "errors"

var queueIntegers []int

func enqueueInteger(val int) {
	queue = append(queue, val)
}

func dequeueInteger(val int) (int, error) { 
	if len(queue) == 0 {
		return -1, errors.New("Queue is empty")

	}
	tmp := queue[0]
	queue = queue[1:len(queue)]
	return tmp, nil
}