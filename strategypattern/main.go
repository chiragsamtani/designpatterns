package main

func main() {
	root := &Node{5, nil, nil}
	tree := Tree{}

	tree.root = root
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(3)
	tree.setTraversalStrategy(&Inorder{})
	tree.Traverse()

	// swap algorithms at runtime

	tree.setTraversalStrategy(&Preorder{})
	tree.Traverse()

	tree.setTraversalStrategy(&Postorder{})
	tree.Traverse()
}
