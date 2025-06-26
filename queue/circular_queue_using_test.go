package queue

import (
	"fmt"
	"math"
	"testing"
)

const empty = math.MinInt64

// [5,-1,-1] h 0 t 0
// [5,3,-1] h 0 t 1
// [5,3,5] h 0 t 2
// [6,3,5] h 1 t 0


type (

	CircularQueue struct {
		queue	[]int
		size	int
		front	int
		rear 	int
	}

)

func NewCircularQueue(size int) *CircularQueue {
	data := make([]int, size)
	for i := range data {
		data[i] = empty
	}
	return &CircularQueue{
		queue: data,
		size: size,
		front: 0,
		rear: -1,
	}
}

func (queue *CircularQueue) enqueue(i int) {
	queue.rear = queue.incrementPosition(queue.rear)

	if queue.queue[queue.rear] != empty {
		queue.front = queue.incrementPosition(queue.front)
	}

	queue.queue[queue.rear] = i
}

func (queue *CircularQueue) dequeue() (int,error) {
	if queue.queue[queue.front] == empty {
		return -1, fmt.Errorf("empty queue")
	}

	var nextFront int = queue.incrementPosition(queue.front)

	var output int = queue.queue[queue.front]
	queue.queue[queue.front] = empty
	queue.front = nextFront

	return output, nil
}

func (queue *CircularQueue) incrementPosition(pos int) int{
	pos++

	if pos == len(queue.queue) {
		return 0
	}

	return pos
}

func TestEnqueueDequeueBasic(t *testing.T) {
	q := NewCircularQueue(3)

	q.enqueue(5)
	val, err := q.dequeue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 5 {
		t.Errorf("expected 5, got %d", val)
	}
}

func TestEnqueueWrapAround(t *testing.T) {
	q := NewCircularQueue(3)

	q.enqueue(5)
	q.enqueue(3)
	q.enqueue(5) // overwrites first value
	q.enqueue(4)
	val, _ := q.dequeue() // should return 3
	if val != 3 {
		t.Errorf("expected 3, got %d", val)
	}
}

func TestOverwriteOldest(t *testing.T) {
	q := NewCircularQueue(3)

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4) // overwrites 1
	val, _ := q.dequeue()
	if val != 2 {
		t.Errorf("expected 2 (since 1 was overwritten), got %d", val)
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := NewCircularQueue(2)

	_, err := q.dequeue()
	if err == nil {
		t.Error("expected error on dequeue from empty queue")
	}
}

func TestFullCycle(t *testing.T) {
	q := NewCircularQueue(2)

	q.enqueue(10)
	q.enqueue(20)
	q.dequeue()
	q.enqueue(30)
	q.dequeue()
	q.enqueue(40)

	expected := []int{30, 40}
	for _, want := range expected {
		got, err := q.dequeue()
		if err != nil || got != want {
			t.Errorf("expected %d, got %d, err: %v", want, got, err)
		}
	}
}