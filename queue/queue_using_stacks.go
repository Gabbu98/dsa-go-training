package queue

import "errors"

type (

	UsingStacks struct {
		stack1 Stack
		stack2 Stack
	}
	
	Stack struct {
		stack []int
	}

)


func (stack *Stack) push(i int) {
	stack.stack = append(stack.stack, i)
}

func (stack *Stack) pop() (int,error) {

	if len(stack.stack) == 0 {
		return -1, errors.New("Stack is empty")
	}

	var i int = stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]

	return i,nil
}

func (usingStack *UsingStacks) enqueue(i int) {
	usingStack.stack1.push(i)
}

func (usingStack *UsingStacks) dequeue() (int,error) {
	i := 0
	for len(usingStack.stack1.stack)!=0 {
		i, err := usingStack.stack1.pop()

		if err!=nil {
			return -1, err
		}
		
		if len(usingStack.stack1.stack)!=0 {
			usingStack.stack2.push(i)
		}
	}

	return i, nil
}


