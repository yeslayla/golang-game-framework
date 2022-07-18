package node

type ProcessMode uint8

const (
	DefaultProcessMode ProcessMode = iota
	PausedProcessMode
	ActiveProcessMode
)

type INode interface {
	GetName() string
	Update() error
	ready(INode) error
	IsProcessing() bool
	GetProcessMode() ProcessMode
	SetProcessMode(ProcessMode)
	Free() error
}
