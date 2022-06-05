package node

type INode interface {
	GetName() string
	Update() error
	ready(INode) error
}
