package demo

import "fmt"

type AvlNode struct {
	Value  int
	Height int
	Left   *AvlNode
	Right  *AvlNode
}

func NewNode(val int) *AvlNode {
	return &AvlNode{Value: val}
}

type AvlTree struct {
	Root *AvlNode
}

// 空节点高度-1，叶子节点高度为0
func (t *AvlTree) height(node *AvlNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}

// 节点的高度为子节点的高度+1
func (t *AvlTree) updateHeight(node *AvlNode) {
	lh := t.height(node.Left)
	rh := t.height(node.Right)
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

// 节点平衡因子：左节点的高度 - 右节点的高度
func (t *AvlTree) blanceFator(node *AvlNode) int {
	if node == nil {
		return 0
	}
	return t.height(node.Left) - t.height(node.Right)
}

// 节点旋转规则
// 节点平衡因子		子节点平衡因子		旋转方式
//
//	> 1				>= 0			右旋
//	> 1				< 0			先左旋再右旋
//	< -1			<= 0			左旋
//	< -1			> 0			先右旋再左旋

func (t *AvlTree) rotate(node *AvlNode) *AvlNode {
	bf := t.blanceFator(node)
	// 左偏树
	if bf > 1 {
		if t.blanceFator(node.Left) >= 0 {
			// 右旋
			return t.rotateRight(node)
		} else {
			// 先左旋再右旋
			node.Left = t.rotateLeft(node)
			return t.rotateRight(node)
		}
	}
	// 右偏树
	if bf < -1 {
		if t.blanceFator(node.Right) <= 0 {
			// 左旋
			return t.rotateLeft(node)
		} else {
			// 先右旋再左旋
			node.Right = t.rotateRight(node)
			return t.rotateLeft(node)
		}
	}
	return node
}

// 右旋规则
// 失衡节点记为node，左子节点为child，child的右子节点记为grandChild
// child作为node节点，将grandChild作为node的左子节点
func (t *AvlTree) rotateRight(node *AvlNode) *AvlNode {
	child := node.Left
	grandChild := child.Right
	// 进行右旋
	child.Right = node
	node.Left = grandChild
	// 更新旋转节点的高度
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

// 左旋规则，和右旋对称
// 失衡节点记为node，右子节点为child，child的左子节点记为grandChild
// child作为node节点，将grandChild作为node的右子节点
func (t *AvlTree) rotateLeft(node *AvlNode) *AvlNode {
	child := node.Right
	grandChild := child.Left
	// 进行左旋
	child.Left = node
	node.Right = grandChild
	// 更新旋转节点的高度
	t.updateHeight(node)
	t.updateHeight(child)
	return child
}

func (t *AvlTree) Insert(value int) {
	t.Root = t.doInsert(t.Root, value)
}

func (t *AvlTree) doInsert(root *AvlNode, value int) *AvlNode {
	if root == nil {
		return NewNode(value)
	}
	if value < root.Value {
		root.Left = t.doInsert(root.Left, value)
	} else if value > root.Value {
		root.Right = t.doInsert(root.Right, value)
	} else {
		return root
	}
	t.updateHeight(root)
	root = t.rotate(root)
	return root
}

func (t *AvlTree) Print() {
	t.doPrint(t.Root)
}

func (t *AvlTree) doPrint(root *AvlNode) {
	if root == nil {
		return
	}
	t.doPrint(root.Left)
	fmt.Printf("%d ", root.Value)
	t.doPrint(root.Right)
}
