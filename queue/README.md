# Queue
- operate on FIFO principle
- (enqueued) added to the end, (dequeued) removed from the front
- another variation = double ended queue -> allows dequeueing from both ends (FIFO / LIFA)

## Implementation
- similar to stacks implemented using double linked lists or arrays and slices

## Complexity
- enqueue and dequeue = O(1)
- linked list faster 
- slcies due to size variation are generally O(n) but in go it is faster

## Application
- queues widely used in graph related problems