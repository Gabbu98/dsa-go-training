package queue

import (
	"reflect"
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

type Tree struct {
	rootNode 	*Node
	leftStack 	[]string
	rightStack	[]string
	isSymmetric bool
}

type Node struct {
	value	string
	left	*Node
	right	*Node
}

func New(values []string) Tree {
	if len(values) == 0 {
		t := Tree{}
		t.isSymmetric = false
		return t
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

	t := Tree{rootNode: root}

	if len(values) % 2 == 0 {
		t.isSymmetric = false
		return t
	} 
	
	t.isSymmetric = t.isSymmetrical()
	return t
}

func (t *Tree) isSymmetrical() bool {
	leftList := []string{}
	rec(*t.rootNode, &leftList, true)
	rightList := []string{}
	rec(*t.rootNode, &rightList, false)

	return reflect.DeepEqual(leftList,rightList)
}

func rec(root Node, list *[]string, left bool) string {

	if left ==true{
		if root.left!=nil {
			*list = append(*list, rec(*root.left,list,left))
			
		}

		if root.right!=nil{
			*list = append(*list, rec(*root.right,list,left))
		}
	} else {
		if root.right!=nil {
			*list = append(*list, rec(*root.right,list,left))
			
		}

		if root.left!=nil{
			*list = append(*list, rec(*root.left,list,left))
		}
	}

	return root.value
}

func TestIsTreeSymmetrical(t *testing.T) {
	tests := []struct {
		tree		[]string
		isSymmetric bool
	}{
		// {[]string{""}, true},
		// {[]string{"1"}, true},
		// {[]string{"1", "2", "2"}, true},
		// {[]string{"1", "2", "3"}, false},
		// {[]string{"1", "2", "2", "3", "", "", "3"}, true},
		{[]string{"1", "2", "", "4"}, false},
		{[]string{"1", "2", "3", "4", "", "5", "6"}, false},
		{[]string{"1", "2", "", "4", "", "5", "6"}, false},
		{[]string{"2", "4", "4", "5", "6", "5", "6"}, false},
		{[]string{"2", "4", "4", "5", "6", "6", "5"}, true},
	}

	for i:=0; i < len(tests); i++{
		tree:=New(tests[i].tree)
		if tree.isSymmetric != tests[i].isSymmetric {
			t.Fatalf("Expected %t but got %t for test case: %d", tests[i].isSymmetric, tree.isSymmetric, i)
		}
	}
}