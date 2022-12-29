package binary_search_tree

/*
* @author: Chen Chiheng
* @date: 2022/12/28 0028 10:21
* @version: 1.0
* @description: 二叉查找树
**/

type Node struct {
	left, right, parent *Node
	Key                 int
}

type BinarySearchTree struct {
	root *Node
}

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Search 在以n为根的树中查找key（递归版本）
func Search(n *Node, key int) *Node {
	current := n
	if current == nil || current.Key == key {
		return current
	}
	if current.Key > key {
		return Search(current.left, key)
	} else {
		return Search(current.right, key)
	}
}

// IterativeSearch 在以n为根的树中查找key（循环版本）
func IterativeSearch(n *Node, key int) *Node {
	current := n
	for current != nil || current.Key != key {
		if key > current.Key {
			current = current.right
		} else {
			current = current.left
		}
	}
	return current
}

// Minimum 查找以n为根的树中key最小的结点
func Minimum(n *Node) *Node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

// Maximum 查找以n为根树中key最大的结点
func Maximum(n *Node) *Node {
	current := n
	for current.right != nil {
		current = current.right
	}
	return current
}

// Predecessor 查找前驱结点
// 前驱结点为比n的key小的结点集合中最大key的结点
func Predecessor(n *Node) *Node {
	current := n
	for current.left != nil {
		Predecessor(current.left)
	}
	pn := current.parent
	for pn != nil && current == pn.left {
		current = pn
		pn = pn.parent
	}
	return pn
}

// Successor 查找后继结点
// 后继结点为大于key[n]结点集合中最小key结点
func Successor(n *Node) *Node {
	current := n // current指针保存当前结点
	// 大于结点current的结点都在current的右子树中
	// 我们先判断current的右子树是否存在
	for current.right != nil { // current的右子树存在，则successor[current]在右子树的最左边
		return Successor(current.left) // 因此在current的左子树查找，同时移动current指针
	}
	// current的右子树不存在，此时若parent[current]不存在，则不存在后继结点
	// 若是左子结点,则其父结点为后继结点，（二叉查找树性质key[left] < key[parent] < key[right]）
	// 若parent[current]存在，我们需要判断current是左子结点还是右子结点
	pn := current.parent                   // 获取当前结点current的父结点
	for pn != nil && current == pn.right { // 判断current的父结点pn是否存在，若存在current是否为右子结点
		current = pn
		pn = pn.parent
	}
	return pn
}

// Insert 插入新的值到二叉查找树中
// 新插入的结点必定为叶子结点
func Insert(n *Node, key int) {
	nn := &Node{Key: key} // 根据key创建新结点
	current := n          // 当前结点指向n，n为树的根结点，可能为nil
	var cpn *Node         // 用于保存当前结点current的父结点
	for current != nil {  // 若current指向的结点不为空，则查找新结点nn是在其右子树还是左子树
		cpn = current             // 保存为父结点
		if nn.Key > current.Key { // 若新节点Key[nn] > key[current]，则新结点nn应该在current的右子树中
			current = current.right
		} else { // 相反则新结点nn应该在current的左子树中
			current = current.left
		}
	}
	// 退出循环后，当前current必定为nil，我们需要将新结点nn放到current所在的位置
	nn.parent = cpn
	// 由此我们需要判断current是其父结点的左子结点还是右子结点
	if nn.Key > cpn.Key {
		cpn.right = nn
	} else {
		cpn.left = nn
	}
}

// Delete 删除等于key的结点
func Delete(n *Node, key int) {

}
