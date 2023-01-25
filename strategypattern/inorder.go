package main

import "fmt"

// concrete strategy

type Inorder struct {
}

func (i *Inorder) traverse(t *Node) {
	if t == nil {
		return
	}
	i.traverse(t.left)
	fmt.Println(t.val)
	i.traverse(t.right)
}
