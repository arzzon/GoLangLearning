package main

import (
	"container/list"
	"fmt"
)

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

// BinarySearchTree represents the BinarySearchTree
type BinarySearchTree struct {
	root *node
	size int
	// Insert
	// delete
	// search
	// inorder,preorder,postorder
}

// NewBinarySearchTree creates a new BinarySearchTree
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
		size: 0,
	}
}

// Insert inserts a new node to the BinarySearchTree
func (t *BinarySearchTree) Insert(data int) {
	if t.root == nil {
		t.root = newNode(data)
		t.size++
	} else {
		tempParentNode := t.root
		tempNode := tempParentNode
		for tempNode != nil {
			if data < tempNode.data {
				tempParentNode = tempNode
				tempNode = tempParentNode.left
			} else if data > tempNode.data {
				tempParentNode = tempNode
				tempNode = tempParentNode.right
			}
		}
		if data < tempParentNode.data {
			tempParentNode.left = newNode(data)
		} else {
			tempParentNode.right = newNode(data)
		}
	}
}

// Remove removes a specific node from the tree
// func (t *BinarySearchTree) Remove(data int) {
// 	if t.root == nil {
// 		return
// 	}
// 	tempNode := t.root
// 	tempParentNode := tempNode
// 	for tempNode != nil && tempNode.data != data {
// 		if data < tempNode.data {
// 			tempParentNode = tempNode
// 			tempNode = tempNode.left
// 		} else {
// 			tempParentNode = tempNode
// 			tempNode = tempNode.right
// 		}
// 	}
// 	if tempNode == nil {
// 		return
// 	}
// 	if tempNode == t.root {
// 		t.root = nil
// 	} else if tempNode.left == nil && tempNode.right == nil {
// 		if tempParentNode.left != nil && tempParentNode.left.data == data {
// 			tempParentNode.left = nil
// 		} else {
// 			tempParentNode.right = nil
// 		}
// 	} else {
// 		targetNodeParent := tempParentNode
// 		for tempNode != nil && tempNode.data != data {
// 			if data < tempNode.data {
// 				tempParentNode = tempNode
// 				tempNode = tempNode.left
// 			} else {
// 				tempParentNode = tempNode
// 				tempNode = tempNode.right
// 			}
// 		}
// 	}

// }

// Preorder traverses BinarySearchTree in VLR order
func (t *BinarySearchTree) Preorder(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	t.Preorder(n.left)
	t.Preorder(n.right)
}

// Inorder traverses BinarySearchTree in LVR order
func (t *BinarySearchTree) Inorder(n *node) {
	if n == nil {
		return
	}
	t.Inorder(n.left)
	fmt.Println(n.data)
	t.Inorder(n.right)
}

// Postorder traverses BinarySearchTree in LRV order
func (t *BinarySearchTree) Postorder(n *node) {
	if n == nil {
		return
	}
	t.Postorder(n.left)
	t.Postorder(n.right)
	fmt.Println(n.data)
}

// LevelOrder traverses the BinarySearchTree in level wise
func (t *BinarySearchTree) LevelOrder() {
	// If tree is empty then return
	if t.root == nil {
		return
	}
	// Create a queue
	q := list.New()
	// Enqueue the root node first
	q.PushBack(t.root)
	// nodeEle stores the first node stored in the qeueue encapsulated inside Element{}
	nodeEle := q.Front()
	// Dequeue till queue is empty
	for nodeEle != nil {
		// Get the node and print the data
		n := nodeEle.Value.(*node)
		fmt.Println(n.data)
		if n.left != nil { // If there is a left child then enqueue it
			q.PushBack(n.left)
		}
		if n.right != nil { // If there is a right child then enqueue it as well
			q.PushBack(n.right)
		}
		q.Remove(nodeEle)   // Dequeue
		nodeEle = q.Front() // Peek for the front element in the queue
	}
}

// // LevelOrder traverses the BinarySearchTree in level wise
// func (t *BinarySearchTree) LevelOrderPretty() {
// 	// If tree is empty then return
// 	if t.root == nil {
// 		return
// 	}
// 	// Create a queue
// 	q := list.New()
// 	// Enqueue the root node first
// 	q.PushBack(t.root)
// 	q.PushBack(nil)
// 	// nodeEle stores the first node stored in the qeueue encapsulated inside Element{}
// 	nodeEle := q.Front()
// 	// Dequeue till queue is empty
// 	for nodeEle != nil {
// 		if nodeEle.Value == nil {
// 			fmt.Println()
// 			q.Remove(nodeEle)
// 			nodeEle = q.Front()
// 			continue
// 		}
// 		// Get the node and print the data
// 		n := nodeEle.Value.(*node)
// 		fmt.Printf("%d ", n.data)
// 		if n.left != nil { // If there is a left child then enqueue it
// 			q.PushBack(n.left)
// 		}
// 		if n.right != nil { // If there is a right child then enqueue it as well
// 			q.PushBack(n.right)
// 			q.Remove(nodeEle) // Dequeue
// 			nodeEle = q.Front()
// 			if q.Front() == nil {
// 				q.PushBack(nil)
// 			}
// 			continue
// 		}
// 		q.Remove(nodeEle)   // Dequeue
// 		nodeEle = q.Front() // Peek for the front element in the queue
// 	}
// }

// Traverse traverses the BinarySearchTree in Preorder fashion
func (t *BinarySearchTree) Traverse(data int) {
	if t.root == nil {
		t.root = newNode(data)
	}
}

func main() {
	tree := NewBinarySearchTree()
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)

	fmt.Println("Preorder")
	tree.Preorder(tree.root)
	fmt.Println("Inorder")
	tree.Inorder(tree.root)
	fmt.Println("Postorder")
	tree.Postorder(tree.root)
	fmt.Println("LevelOrder")
	tree.LevelOrder()
	fmt.Println()
	//tree.LevelOrderPretty()
}
