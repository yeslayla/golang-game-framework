package rendering

import "github.com/manleydev/golang-game-framework/core"

type Camera2D interface {
	GetZoom() float64
	SetZoom(float64)

	GetOffset() core.Vector2
	SetOffset(core.Vector2)

	GetModifiers() Camera2DModifiers
}

type Camera2DModifiers struct {
	Zoom   float64
	Offset core.Vector2
}
