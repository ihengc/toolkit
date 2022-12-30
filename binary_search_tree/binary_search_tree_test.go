package binary_search_tree

import "testing"

/*
* @author: Chen Chiheng
* @date: 2022/12/28 0028 14:51
* @version: 1.0
* @description:
**/

func TestRecursiveMinimum(t *testing.T) {
	bst := New()
	bst.root = &Node{Key: 10}
	bst.root.left = &Node{Key: 5}
	bst.root.left.right = &Node{Key: 7}
	bst.root.left.left = &Node{Key: 4}
	bst.root.right = &Node{Key: 20}
	testRecursiveMinimum(bst.root, 4, t)
}

func testRecursiveMinimum(n *Node, minimum int, t *testing.T) {
	minN := RecursiveMinimum(n)
	if minN.Key != minimum {
		t.Fatal("RecursiveMinimum Error")
	}
}
