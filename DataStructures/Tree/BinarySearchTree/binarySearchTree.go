package main

import (
	"container/list"
	"fmt"
)

// node represents a single node in the tree
type node struct {
	//parent *node
	data  int
	left  *node
	right *node
}

// newNode creates a node
func newNode(data int) *node {
	return &node{
		//parent: nil,
		data:  data,
		left:  nil,
		right: nil,
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
func (t *BinarySearchTree) Remove(root *node, data int) *node {
	if root == nil {
		return nil
	}
	if root.left != nil && data < root.data {
		root.left = t.Remove(root.left, data)
	} else if root.right != nil && data > root.data {
		root.right = t.Remove(root.right, data)
	} else {
		if root.left == nil && root.right == nil { // Leaf node to be removed
			root = nil
			//return root
		} else if root.left == nil { // right node to be removed(node with single child).
			root = root.right
			//return root
		} else if root.right == nil { // left node to be removed(node with single child).
			root = root.left
			//return root
		} else { // node with 2 children
			tempNode := t.FindMinNode(root.right)
			// store the data of the node that is going to be replaced.
			data = tempNode.data
			root = t.Remove(root, tempNode.data)
			root.data = data
		}
	}
	return root

}

func (t *BinarySearchTree) FindMinNode(root *node) *node {
	if root == nil {
		return nil
	}
	tempNode := root
	tempParentNode := tempNode
	for tempNode != nil {
		tempParentNode = tempNode
		tempNode = tempNode.left

	}
	return tempParentNode
}

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
	fmt.Println("Remove")
	tree.root = tree.Remove(tree.root, 2)
	fmt.Println("LevelOrder")
	tree.LevelOrder()
	//tree.LevelOrderPretty()
}

/*
			4
		2		6
	  1	  3   5   7
*/
