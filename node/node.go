package node

import (
	"fmt"
	"log"

	"github.com/manleydev/golang-game-framework/rendering"
)

type Node struct {
	Name        string
	children    []INode
	parent      INode
	processMode ProcessMode

	onReadyMethods  []func() error
	onUpdateMethods []func() error
	onDraw2dMethods []func(rendering.Renderer2D) error
}

func (node *Node) ready(parent INode) error {
	node.parent = parent

	for _, readyMethod := range node.onReadyMethods {
		if err := readyMethod(); err != nil {
			return err
		}
	}
	return nil
}

func (node *Node) OnReady(callback func() error) {
	if callback == nil {
		return
	}
	node.onReadyMethods = append(node.onReadyMethods, callback)
}

func (node *Node) Update() error {
	for _, child := range node.children {
		if err := child.Update(); err != nil {
			return err
		}
	}

	if node.IsProcessing() {
		for _, updateMethod := range node.onUpdateMethods {
			if err := updateMethod(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (node *Node) OnUpdate(callback func() error) {
	if callback == nil {
		return
	}
	node.onUpdateMethods = append(node.onUpdateMethods, callback)
}

func (node *Node) Draw(renderer rendering.Renderer2D) error {
	for _, child := range node.children {
		drawable, ok := child.(interface {
			Draw(rendering.Renderer2D) error
		})
		if !ok {
			continue
		}
		if err := drawable.Draw(renderer); err != nil {
			return err
		}
	}

	for _, drawMethod := range node.onDraw2dMethods {
		if err := drawMethod(renderer); err != nil {
			return err
		}
	}

	return nil
}

func (node *Node) OnDraw2D(callback func(rendering.Renderer2D) error) {
	if callback == nil {
		return
	}
	node.onDraw2dMethods = append(node.onDraw2dMethods, callback)
}

func (node *Node) IsProcessing() bool {
	if node.processMode == DefaultProcessMode {
		if node.parent != nil {
			return node.parent.IsProcessing()
		}
		return true
	}

	return node.processMode == ActiveProcessMode
}

func (node *Node) GetProcessMode() ProcessMode {
	return node.processMode
}

func (node *Node) SetProcessMode(mode ProcessMode) {
	node.processMode = mode
}

func (node *Node) GetName() string {
	return node.Name
}

func (node *Node) Free() error {

	var parent *Node = node.parent.(*Node)
	if err := parent.removeChild(node); err != nil {
		return err
	}
	node.parent = nil
	return nil
}

func (node *Node) removeChild(target *Node) error {
	for i, child := range node.children {
		if child == target {
			node.children = append(node.children[:i], node.children[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not remove child '%s', child not found on parent!", target.Name)
}

func (node *Node) GetChild(index int) INode {
	if index < len(node.children) {
		return nil
	}
	return node.children[index]
}

func (node *Node) internalAddChild(parent INode, child INode) {
	node.children = append(node.children, child)

	if err := child.ready(parent); err != nil {
		log.Fatalf("Node(%s) AddChild: %v", node.Name, err)
	}
}

func (node *Node) AddChild(child INode) {
	node.internalAddChild(node, child)
}

func (node *Node) GetParent() INode {
	return node.parent
}

func NewNode() Node {
	return Node{
		children:    []INode{},
		processMode: DefaultProcessMode,
	}
}
