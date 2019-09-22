package main

import "fmt"

// node represents a single node in the tree
type node struct {
	parent *node
	data   int
	left   *node
	right  *node
}

// newNode creates a node
func newNode(data int) *node {
	return &node{
		parent: nil,
		data:   data,
		left:   nil,
		right:  nil,
	}
}

// BinaryTree represents the BinaryTree
type BinaryTree struct {
	root *node
	size int
	// Insert
	// delete
	// search
	// inorder,preorder,postorder
}

// NewBinaryTree creates a new BinaryTree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
		size: 0,
	}
}

// Insert inserts a new node to the BinaryTree
func (t *BinaryTree) Insert(data int) {
	// If tree is empty
	if t.root == nil {
		t.root = newNode(data)
		t.size++
	} else {
		tempParentNode := t.root
		tempNode := tempParentNode
		for tempNode != nil {
			if tempNode.left == nil { // Start from left
				tempNode.left = newNode(data)
				t.size++
				break
			} else if tempNode.right == nil { // If left is filled try for right
				tempNode.right = newNode(data)
				t.size++
				break
			} else { // If both nodes are filled then move to left node
				tempParentNode = tempNode
				tempNode = tempParentNode.left
			}
		}
	}
}

// Remove removes a specific node from the tree
func (t *BinaryTree) Remove(data int) {

}

// Preorder traverses BinaryTree in VLR order
func (t *BinaryTree) Preorder(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	t.Preorder(n.left)
	t.Preorder(n.right)
}

// Inorder traverses BinaryTree in LVR order
func (t *BinaryTree) Inorder(n *node) {
	if n == nil {
		return
	}
	t.Inorder(n.left)
	fmt.Println(n.data)
	t.Inorder(n.right)
}

// Postorder traverses BinaryTree in LRV order
func (t *BinaryTree) Postorder(n *node) {
	if n == nil {
		return
	}
	t.Postorder(n.left)
	t.Postorder(n.right)
	fmt.Println(n.data)
}

// Traverse traverses the BinaryTree in Preorder fashion
func (t *BinaryTree) Traverse(data int) {
	t.Preorder(t.root)
}

func main() {
	tree := NewBinaryTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Preorder(tree.root)
	fmt.Println()
	tree.Inorder(tree.root)
	fmt.Println()
	tree.Postorder(tree.root)
}
