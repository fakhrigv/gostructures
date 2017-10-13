package rbtree

import (
	"fmt"
	"log"
)

type color bool

const (
	red   color = true
	black color = false
)

var (
	nilNode = &RBTree{0, 0, black, nil, nil, nil}
)

// RBTree ...
type RBTree struct {
	key   uint32
	value interface{}

	color color // true - red, false - black

	left  *RBTree
	right *RBTree

	parent *RBTree
}

// NewRBTree ...
func NewRBTree() *RBTree {
	tree := &RBTree{
		color:  black,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	return tree
}

// AddNode ...
func (tree *RBTree) AddNode(key uint32, value interface{}) {
	fmt.Println("AddNode:", key, value, "tree:", tree)
	if tree.left == nil {
		tree.key = key
		tree.value = value

		tree.color = black

		tree.left = nilNode
		tree.right = nilNode

		tree.parent = nil

		return
	}

	tree.addNode(key, value)

	tree.color = black
}

func (tree *RBTree) addNode(key uint32, value interface{}) {
	var cTree *RBTree

	if tree.key > key {
		// move left
		cTree = tree.left
	} else if tree.key < key {
		// move right
		cTree = tree.right
	} else {
		// there is already this key. Later implement error
		return
	}

	if cTree == nilNode {
		newNode := &RBTree{
			key:    key,
			value:  value,
			color:  red,
			parent: tree,
			left:   nilNode,
			right:  nilNode,
		}

		if tree.key > key {
			tree.left = newNode
		} else {
			tree.right = newNode
		}

		if newNode.parent.color == red {
			log.Println("balancing")
			newNode.rebalance()
		}

		return
	}

	cTree.addNode(key, value)
}

func (tree *RBTree) rebalance() {
	var uncle *RBTree
	var grand *RBTree
	parent := tree.parent

	if tree.parent.parent != nil {
		grand = tree.parent.parent
	} else {
		return
	}

	currentPosition := "left"

	if tree == parent.right {
		currentPosition = "right"
	}

	// case 1: unclePosition == right, uncleColor == red, currentPosition = left | right
	// case 2: unclePosition == right, uncleColor == black, currentPosition = right
	// case 3: unclePosition == right, uncleColor == black, currentPosition = left
	// case 4: unclePosition == left, uncleColor == red, currentPosition = right | left
	// case 5: unclePosition == left, uncleColor == black, currentPosition = left
	// case 6: unclePosition == left, uncleColor == black, currentPosition = right

	if parent == grand.right {
		// uncle position is left
		uncle = grand.left
		if uncle.color == black {
			if currentPosition == "left" {
				// turn parent tree to right
				parent.turnRight()
				tree.rebalance()
			} else {
				// change colors
				parent.color = black
				grand.color = red

				// turn grand tree to left
				grand.turnLeft()
			}
		} else {
			// change colors
			parent.color = black
			uncle.color = black
			grand.color = red
		}
	} else {
		// uncle position is right
		uncle = grand.right
		if uncle.color == black {
			if currentPosition == "right" {
				// turn parent tree to left
				parent.turnLeft()
				tree.rebalance()
			} else {
				// change colors
				parent.color = black
				grand.color = red

				// turn grand tree to right
				grand.turnRight()
			}
		} else {
			// change colors
			parent.color = black
			uncle.color = black
			grand.color = red
		}
	}
}

func (tree *RBTree) turnLeft() {
	var temp RBTree
	temp = *tree

	rTree := temp.right
	clTree := rTree.left

	rTree.left = &temp
	rTree.parent = temp.parent

	temp.left = clTree
	temp.parent = rTree
	temp.right = nilNode

	*tree = *rTree
}

func (tree *RBTree) turnRight() {
	var temp RBTree
	temp = *tree

	lTree := temp.left
	crTree := lTree.right

	lTree.right = &temp
	lTree.parent = temp.parent

	temp.right = crTree
	temp.parent = lTree
	temp.left = nilNode

	*tree = *lTree
}
