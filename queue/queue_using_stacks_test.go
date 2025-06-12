package queue

import (
	"errors"
	"testing"
)

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
	var err error
	for len(usingStack.stack1.stack)!=0 {
		i, err = usingStack.stack1.pop()

		if err!=nil {
			return -1, err
		}
		
		if len(usingStack.stack1.stack)!=0 {
			usingStack.stack2.push(i)
		} 
	}

	for len(usingStack.stack2.stack)!=0 {
		i, err := usingStack.stack2.pop()

		if err!=nil {
			return -1, err
		}
		
		
		usingStack.stack1.push(i)
		
	}

	return i, nil
}


func TestQueueUsingStacks(t *testing.T) {
	tests := []struct {
		enqueue					[]int
		dequeueTimes			int
		expectedLastDequeueItem int
		expectErr				bool
	}{
		{[]int{1,2,3,4},1,1,false},
		{[]int{1,2,3,4},2,2,false},
		{[]int{1,2,3,4},3,3,false},
		{[]int{1,2,3,4},4,4,false},
	}

	for i, test := range tests {
		queue := new(UsingStacks)
		for _, n := range test.enqueue {
			queue.enqueue(n)
		}

		var n int
		var err error
		for j := 1; j <= test.dequeueTimes; j++ {
			n,_ = queue.dequeue()
			if !test.expectErr && err!=nil {
				t.Fatalf("Failed test case #%d. Did not expect an error. Error %s", i, err)
			}
		}
		if n != test.expectedLastDequeueItem {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.expectedLastDequeueItem, n)
		}
	}
}