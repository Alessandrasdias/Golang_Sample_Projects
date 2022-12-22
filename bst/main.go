package main

import "fmt"

// The concept of binary tree is that:
// 1 Everything that goes to the left side is smaller than the parent
// 2 Everything the goes to the right is larger than the parent
// Which makes it excellent for searching
// Also a parent node can only have 2 children
// A node without children is a leaf node
// I will use recursion to build this BST

// First define the struct for the node
// The nodes will represent the components of the binary tree
// I need to have a key that will be used to compare to the other nodes, a pointer to the left side and a pointer to the right side
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Create an Insert method to add nodes to the tree
// The idea is to compare the key to the receiver and go to the left/right child
// if K is larger than my node key move to the right
// If my K is smaller than my node key than move to the left
// Also if My left/right child is empty, I'll have to create a new node and add the value of K

// The first time there will be no nodes and the internal if will suffice
// However, the second time forward, there will be nodes
// So, to do the process of inserting the lager/Smaller I'll use recursion and call the func inside of itself
// It won't be a eternal loop bc I have a base case, which prevents the forever loop
// I'm calling it on the child and not the parent
//Insert a node
func (n *Node) Insert(k int) {
	if k > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if k < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Create an search method
// Returns if the k value was found in the tree or not
// The idea is the same af before here: to compare the key to the receiver and go to the left/right child
// if K is larger than my node key move to the right
// If my K is smaller than my node key than move to the left
// the difference is that if the k value is not larger or smaller we will return true bc we are just checking if it's there
// I will use a recursive call to have the method keep calling itself
// and so that won't be endless we will have base case. If n is nill we got to the end and found not match
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if k > n.Key {
		return n.Right.Search(k)
	} else if k < n.Key {
		return n.Left.Search(k)
	}
	return true
}

func main() {

	//1 st testing the node struct
	// Left and right nodes will show as nil bc I didn't add anything yet
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(44)
	tree.Insert(10)
	tree.Insert(500)
	tree.Insert(36)
	fmt.Println(tree)             // the root will show 100, but its nodes will show the address in memory
	fmt.Println(tree.Search(300)) // true
	fmt.Println(tree.Search(44))  // true
	fmt.Println(tree.Search(22))  // false
	fmt.Println(tree.Search(1))   // false
}
