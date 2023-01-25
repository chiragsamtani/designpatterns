package main

import "fmt"

// concrete strategy

type Postorder struct {
}

func (p *Postorder) traverse(t *Node) {
	if t == nil {
		return
	}
	p.traverse(t.left)
	fmt.Println(t.val)
	p.traverse(t.right)
}
