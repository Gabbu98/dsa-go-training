package queue

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

func (queue *CircularQueue) enqueue(i int) {
	var currentRear int = queue.rear

	queue.rear++

	if queue.rear == len(queue.queue) {
		queue.rear = 0
	}

	queue.queue[currentRear] = i
}

func (queue *CircularQueue) dequeue() int {

	var currentFront int = queue.front

	queue.front++

	if queue.front == len(queue.queue) {
		queue.front = 0
	}

	return queue.queue[currentFront]
	
}