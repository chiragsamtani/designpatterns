package main

import "fmt"

// concrete strategy

type Preorder struct {
}

func (p *Preorder) traverse(t *Node) {
	if t == nil {
		return
	}
	fmt.Println(t.val)
	p.traverse(t.left)
	p.traverse(t.right)
}
