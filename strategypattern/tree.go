package main

// tree.go: acts as context class

type Node struct {
	val   int
	left  *Node
	right *Node
}

// Tree acts as the context class
type Tree struct {
	root     *Node
	strategy TraversalStrategy
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{data, nil, nil}
	} else {
		t.root.Insert(data)
	}
}

func (t *Node) Insert(data int) {
	if data <= t.val {
		if t.left == nil {
			t.left = &Node{data, nil, nil}
		} else {
			t.left.Insert(data)
		}
	} else {
		if t.right == nil {
			t.right = &Node{data, nil, nil}
		} else {
			t.right.Insert(data)
		}
	}
}

// Traverse strategy execution function
func (t *Tree) Traverse() {
	t.strategy.traverse(t.root)
}

// setTraversalStrategy in a context class, we need to have a setter to
// replace the strategy at runtime
func (t *Tree) setTraversalStrategy(ts TraversalStrategy) {
	t.strategy = ts
}
