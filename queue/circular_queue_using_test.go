package queue

import "errors"

// [-1,-1,-1,-1] front 4 rear 4
// en 1 [-1,,-1,-1,1] front 4 rear 0
// en 2 [2,-1,-1,1] front 4 rear 1
// en 3 [2,3,-1,1] front 4 rear 2
// deq  [2,3,-1,-1] front 0 rear 2

type (

	CircularQueue struct {
		queue	[4]int
		front	int
		rear 	int
	}

)

func (queue *CircularQueue) enqueue(i int) error {

}

func (queue *CircularQueue) dequeue() (int,error) {

}