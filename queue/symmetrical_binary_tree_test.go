package queue

import "testing"

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

// type Node struct {
// 	value	int
// 	left	*Node
// 	right	*Node
// }

// func (node *Node) inorder_traversal(root *Node) ([]string, error) {
// 	output := make([]string, 0)
// 	if root == nil {
// 		return output
// 	}

// 	output = append(output, inorder_traversal(root.left)...)
// 	output = append(output, root.value)
// 	output = append(output, inorder_traversal(root.right)...)
// 	return output, nil
// }


func TestIsTreeSymmetrical(t *testing.T) {
	tests := []struct {
		tree		string
		isSymmetric bool
	}{
		{"", false},
		{"1", true},
		{"1,2,2", true},
		{"1,2,3", false},
		{"1,2,2,3,nil,nil,3",3},
	}

	for i, test := range tests{
		if got, err := isTreeSymmetrical(tree.New(test.tree)); err != nil {
			t.Fatalf("failed test case #%d, unexpected error %s", i, err)
		} else if got != test.isSymmetric {
			t.Fatalf("Failed test case #%d got %t", i, test.isSymmetric, got)
		}
	}
},