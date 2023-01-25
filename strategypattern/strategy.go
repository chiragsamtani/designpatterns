package main

type TraversalStrategy interface {
	traverse(t *Node)
}
