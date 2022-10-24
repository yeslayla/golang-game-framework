package rendering

import (
	"github.com/manleydev/golang-game-framework/core"
)

type Renderer2D interface {
	DrawTexture2D(DrawTexture2DInput) error

	Draw() error
	Update(delta float64) error
	Destroy()
	SetCamera(*Camera2D) error
}

type DrawTexture2DInput struct {
	Texture  Texture2D
	Rect     core.Rect2D
	Position core.Vector2
	Rotation float64
}
