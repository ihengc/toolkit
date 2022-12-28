package redblack_tree

/*
* @author: Chen Chiheng
* @date: 2022/12/27 0027 11:53
* @version: 1.0
* @description: 红黑树
**/

// 叶子结点 NIL
// 根叶黑，红父子黑，子孙路径黑相同。
// 黑高度

type RedBlackTree struct {
	root *node
}

func New() *RedBlackTree {
	return &RedBlackTree{}
}

type node struct {
	left, right, parent *node
	color               byte
	value               interface{}
}

// leftRotate 对以x为根的子树进行左旋转
//
//		X							Y
//	a		Y			->		X		c
//		 b		c			a		b
//
// 分为三部分：1.x的右结点
// 2.y的父结点
// 3.y的左结点
func leftRotate(tree *RedBlackTree, x *node) {
	y := x.right
	x.right = y.left
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x.parent.left == x {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func rightRotate(tree *RedBlackTree, x *node) {
	y := x.left
	y.left = x.right
	x.parent = y.parent
	if y.parent == nil {
		tree.root = x
	} else if y.parent.left == y {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

func (tree *RedBlackTree) Insert(value interface{}) {

}
