package queue

import (
	"fmt"
	"testing"
)

//  /*
// TestIsTreeSymmetrical tests solution(s) with the following signature and problem description:

// 	func IsTreeSymmetrical(root *tree.BinaryTreeNode) (bool, error)

// Given a binary tree return true of it is symmetric and false otherwise. A tree is symmetric if you
// can draw a vertical line through the root and then the left subtree is the mirror image of the right subtree.

// 	  Symmetric       Not Symmetric
// 	      2                2
// 	    /   \            /   \
// 	   /     \          /     \
// 	  4       4        3       4
// 	 / \     / \      / \     / \
// 	5   6   6   5    5   6   6   5

// For example given "2,4,4,5,6,6,5", shown in the symmetric tree above return true.
// Given "2,3,4,5,6,6,5", shown in the not symmetric tree above return false.
// */

// // if not odd numbear != symmetrical

// 5642
// 5642

// 5632
// 5642

type Tree struct {
	rootNode 	*Node
	leftStack 	[]string
	rightStack	[]string
}

type Node struct {
	value	string
	left	*Node
	right	*Node
}

func New(values []string) Tree {
	if len(values) == 0 {
		return Tree{}
	}

	i:=1
	root := &Node{value: values[0]}
	queue := []*Node{root}

	for i < len(values) {
		current := queue[0]
		queue = queue[1:]

		if current.left == nil {
			current.left = &Node{}
			current.left.value = values[i]
			queue = append(queue, current.left)
			i++
		}

		if i >= len(values) {
			break
		}

		if current.right == nil {
			current.right = &Node{}
			current.right.value = values[i]
			queue = append(queue, current.right)
			i++
		}
	}

	return Tree{rootNode: root}
}

func TestIsTreeSymmetrical(t *testing.T) {
	tests := []struct {
		tree		[]string
		isSymmetric bool
	}{
		{[]string{"1","2","2","3","","","3"},true},
	}

	for i:=0; i < len(tests); i++{
		tree:=New(tests[i].tree)
		fmt.Print("%w",tree)
		// if got, err := isTreeSymmetrical(tree.New(test.tree)); err != nil {
		// 	t.Fatalf("failed test case #%d, unexpected error %s", i, err)
		// } else if got != test.isSymmetric {
		// 	t.Fatalf("Failed test case #%d got %t", i, test.isSymmetric, got)
		// }
	}
}