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
	/*
		ALGO:
			0. If tree is empty then create a new node and assign that to the root, Else do the following.
			1. Store the root in temporary parent node which will store the parent node for the temp node in each step of iteration.
			2. Store the root in another temporary node, that will tell us where to insert data.
			3. Repeat till temp node is nil
				3.1. if data is less then,
					3.1.1. First keep update the temporary parent node with the current node's address.
					3.1.2. Update the temporary node with it's left child's address.
				3.2. if data is greater then,
					3.2.1. First keep update the temporary parent node with the current node's address.
					3.2.2. Update the temporary node with it's right child's address.
			4. The loop exits beacuse temporary node is nil, that means by using the temporary parent node we can insert the new node.
				4.1. Check if the new node's data is less than the temporary parent node's data, if thats the case, then
					4.1.1. Create the new node and assign the address to the left link of temporary parent node.
				4.2. Else, Create the new node and assign the address to the left link of temporary parent node.
				update size.
	*/
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
		t.size++
	}
}

// Remove removes a specific node from the tree
func (t *BinarySearchTree) Remove(root *node, data int) {
	t.root = t.RemoveNode(root, data)
}

// RemoveNode removes a specific node from the tree
func (t *BinarySearchTree) RemoveNode(root *node, data int) *node {
	/*
		ALGO: RemovNode(root, data)
			[First search for the node to remove]
			1. If root is nil then its the terminating condition of recursion
			2. If there is a left child and data is less than current node's data then data removed has to be in the left subtree,
				so recursively call RemoveNode with the left child node and data and store the returned result in current node's left address.
			3. If there is a right child and data is greater than current node's data then data removed has to be in the right subtree,
				so recursively call RemoveNode with the right child node and data and store the returned result in current node's right address.
			4. Else we have reached the target node.
			[Remove the node]
				Case 1: (Leaf node)
					4.1. return nil that will finally get assigned to its parent node's left/right link in recursive call back.
				Case 2: (node with single child)
					5.2. return that single child's address to the parent node's left/right link in recursive call back.
				Case 3: (node with two children)
					5.3. Find the node with the minimum value in the right subtree
					5.4. Store the minimum node's data in a variable
					5.5. Call the RemoveNode with the minimum value and current node's right child's address and store the returned reference to the current node's right link.
					5.6. Now copy the minimum value in the current node's data
				Update size.
	*/
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
		t.size--
	}
	return root

}

// FindMinNode finds the node with the minimum value
func (t *BinarySearchTree) FindMinNode(root *node) *node {
	/*
		ALGO:
			1. Recurcively/ Iteretively go to the left most element in the tree.
			2. Return it.
	*/
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
