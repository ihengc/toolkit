package redblack_tree

import "testing"

/*
* @author: Chen Chiheng
* @date: 2022/12/27 0027 16:20
* @version: 1.0
* @description:
**/

func TestRedBlackTree(t *testing.T) {
	tree := New()
	x := &node{value: "x"}
	y := &node{value: "y"}
	a := &node{value: "a"}
	x.left = a
	a.parent = x
	b := &node{value: "b"}
	y.left = b
	b.parent = y
	c := &node{value: "c"}
	y.right = c
	c.parent = y
	x.right = y
	y.parent = x
	tree.root = x
	leftRotate(tree, x)
	testLeftRotate(tree, x, y, t)
}

func testLeftRotate(tree *RedBlackTree, x *node, y *node, t *testing.T) {
	if tree.root != y {
		t.Fatal("LeftRotate Error")
	}
	if y.right.value != "c" {
		t.Fatal("LeftRotate Error")
	}
	if y.left != x || x.parent != y {
		t.Fatal("LeftRotate Error")
	}
	if x.left.value != "a" || x.right.value != "b" {
		t.Fatal("LeftRotate Error")
	}
}
