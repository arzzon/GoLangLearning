package main

import (
	"container/list"
	"fmt"
	"reflect"
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
	// remove
	// search
	// inorder,preorder,postorder,level order
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
func (t *BinarySearchTree) Remove(root *node, data int) {
	t.root = t.RemoveNode(root, data)
}

// RemoveNode removes a specific node from the tree
func (t *BinarySearchTree) RemoveNode(root *node, data int) *node {
	if root == nil {
		return nil
	}
	if root.left != nil && data < root.data {
		root.left = t.RemoveNode(root.left, data)
	} else if root.right != nil && data > root.data {
		root.right = t.RemoveNode(root.right, data)
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
			root = t.RemoveNode(root, tempNode.data)
			root.data = data
		}
	}
	return root

}

// FindMinNode finds the node with the minimum value
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
	/*
		ALGO:
				1. First insert root and newLine in queue
				2. Dequeue till queue is empty
					2.1. If dequeued element is a node,then
						2.1.1. Enqueue its left and then right children(left - > right order) if they exist
		NOTE: Take an example and then walk through it.
	*/
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

// LevelOrderPretty traverses the BinarySearchTree in level wise
func (t *BinarySearchTree) LevelOrderPretty() {
	/*
		ALGO:
			1. First insert root and newLine in queue
			2. Dequeue till queue is empty
				2.1. If dequeued element is a node,then
					2.1.2. Enqueue its left and then right children(left - > right order) if they exist
				2.2. If dequeued element is newLine, then
					2.2.1. Check if it's the last newLine, if so then stop
					2.2.2. Else move to next line and enqueue another newLine to the queue.
		NOTE: Take an example and then walk through it.
	*/
	// If tree is empty then return
	if t.root == nil {
		return
	}
	newLine := true
	// Create a queue
	q := list.New()
	// Enqueue the root node first
	q.PushBack(t.root)
	q.PushBack(newLine)
	// nodeEle stores the first node stored in the qeueue encapsulated inside Element{}
	nodeEle := q.Front()
	// Dequeue till queue is empty
	for nodeEle != nil {
		if reflect.TypeOf(nodeEle.Value) == reflect.TypeOf(newLine) {
			fmt.Println()
			q.Remove(nodeEle)
			// If its the last newLine then stop
			if q.Front() == nil {
				break
			}
			q.PushBack(newLine)
		} else {
			n := nodeEle.Value.(*node)
			fmt.Printf("%v\t", n.data)
			q.Remove(nodeEle)
			if n.left != nil {
				q.PushBack(n.left)
			}
			if n.right != nil {
				q.PushBack(n.right)
			}
		}
		nodeEle = q.Front()
	}
}

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
	tree.Remove(tree.root, 2)
	fmt.Println("Pretty Level order")
	tree.LevelOrderPretty()
}

/*
			4
		2		6
	  1	  3   5   7
*/
