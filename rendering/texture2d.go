package rendering

import "github.com/manleydev/golang-game-framework/core"

type Texture2D interface {
	Destroy()
	GetCenter() core.Vector2
	GetRect() core.Rect2D
}
