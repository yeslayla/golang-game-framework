package rendering

import "github.com/manleydev/golang-game-framework/core"

type StandardCamera2D struct {
	zoom   float64
	offset core.Vector2
}

func (camera *StandardCamera2D) GetZoom() float64 {
	return camera.zoom
}
func (camera *StandardCamera2D) SetZoom(zoom float64) {
	camera.zoom = zoom
}

func (camera *StandardCamera2D) GetOffset() core.Vector2 {
	return camera.offset
}
func (camera *StandardCamera2D) SetOffset(offset core.Vector2) {
	camera.offset = offset
}

func (camera *StandardCamera2D) GetModifiers() Camera2DModifiers {
	return Camera2DModifiers{
		Zoom:   camera.zoom,
		Offset: camera.offset,
	}
}

func NewStandardCamera() *Camera2D {
	var camera Camera2D = &StandardCamera2D{
		zoom:   1,
		offset: core.Vector2{},
	}
	return &camera
}
