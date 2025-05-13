# Stack
- A data structure
- Operates on LIFO (Last in first out) principle.
- Operations allowed
    - push
    - pop
    - top operations:
        - add
        - remove
        - read
- like a backpack, items placed and removed in reverse order

## Implementation in GO
- via linked lists or arrays and slices
```go
package main

import(
    "container/list"
    "errors"
)

var stack = list.New()

func main() {
    push(1)
    push(2)
    push(3)
    pop()
    pop()
    pop()
}

func push(val int) {
    stack.PushBack(val)
}

func pop() (int, error) {
    if stack.Len() == 0 {
        return -1, errors.New("stack is empty")
    }
    return stack,Remove(stack.Back()).(int), nil
}
```

## Complexity
- Push and Pop = O(1) operations

## Application
- Memory
    - Stack and heap
        - stack used for function calling