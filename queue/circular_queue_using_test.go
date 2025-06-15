package queue

import "testing"

// [-1,-1,-1,-1] front 4 rear 4
// en 1 [-1,,-1,-1,1] front 4 rear 0
// en 2 [2,-1,-1,1] front 4 rear 1
// en 3 [2,3,-1,1] front 4 rear 2
// deq  [2,3,-1,-1] front 0 rear 2

type (
	CircularQueue struct {
		queue	[]int
		front	int
		rear 	int
	}

)

func NewCircularQueue(size int) (*CircularQueue) {
	circQueue := &CircularQueue{
		queue: make([]int, size),
		front: size-1,
		rear: size-1,
	}

	return circQueue
}

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

func TestCircularQueue(t *testing.T) {
	type testRound struct {
		enqueueStart int
		enqueueEnd int
		dequeueStart int
		dequeueEnd int
		expectedErr error
	}

	tests := []struct{
		size int
		testRounds []testRound
	}{
		{1, []testRound{{1, 6, 6, 6, nil}}},
		{2, []testRound{{1, 6, 5, 5, nil}}},
		{4, []testRound{{1, 6, 3, 6, nil}}},
		{4, []testRound{{1, 6, 3, 6, nil}, {1, 6, 3, 6, nil}}},
	}

	for i, test := range tests {
		queue := NewCircularQueue(test.size)

		for j, testRound := range test.testRounds {
			for i := testRound.enqueueStart; i <= testRound.enqueueEnd; i++ {
				queue.enqueue(i)
			}

			for want := testRound.dequeueStart; want <= testRound.dequeueEnd; want++ {
				got, err := queue.dequeue()
				if err != nil {
					if err != testRound.expectedErr {
						t.Fatalf("Failed test case #%d round #%d. Unexpected error %s", i, j, err)
					}
					break
				}

				if got != want {
					t.Fatalf("Failed test case #%d round #%d. Want %d, got %d", i, j, want, got)
				}
			}
		}

	}
}