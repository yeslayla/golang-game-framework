package node

import (
	"github.com/manleydev/golang-game-framework/core"
)

type Node2D struct {
	Node
	Position core.Vector2
	Rotation float64
}

func (node *Node2D) GetPosition() core.Vector2 {
	return node.Position
}

func (node *Node2D) GetGlobalPosition() core.Vector2 {
	if node.parent == nil {
		return node.Position
	}
	global2d, ok := (node.parent).(interface {
		GetGlobalPosition() core.Vector2
	})
	if ok {
		return node.Position.Add(global2d.GetGlobalPosition())
	}
	return node.Position
}

func (node *Node2D) AddChild(child INode) {
	node.internalAddChild(node, child)
}

func NewNode2D() Node2D {
	return Node2D{
		Node: NewNode(),
	}
}
