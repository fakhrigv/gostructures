package rbtree

import "testing"

func TestRBTreeAddNode(t *testing.T) {
	tree := NewRBTree()

	tree.AddNode(5, "Test 5")
	tree.AddNode(4, "Test 4")
	tree.AddNode(3, "Test 3")
	tree.AddNode(2, "Test 2")

	if tree.key != 4 {
		t.Errorf("Failed to add node to tree. Expected root key = 4, instead key = %d", tree.key)
	}

	if tree.left.key != 3 {
		t.Errorf("Failed to add node to tree. Expected root left node key = 4, instead key = %d", tree.left.key)
	}

	if tree.right.key != 5 {
		t.Errorf("Failed to add node to tree. Expected root right node key = 4, instead key = %d", tree.right.key)
	}

	if tree.color != black {
		t.Errorf("Failed to add node to tree. Expected root color = black(false), instead key = %t", tree.color)
	}

	if tree.left.color != black {
		t.Errorf("Failed to add node to tree. Expected root left node color = black(false), instead key = %t", tree.left.color)
	}

	if tree.right.color != black {
		t.Errorf("Failed to add node to tree. Expected root right node color = black(false), instead key = %t", tree.right.color)
	}

	if tree.left.left.color != red {
		t.Errorf("Failed to add node to tree. Expected root left left node color = red(true), instead key = %t", tree.left.left.color)
	}
}
