package node

import (
	"errors"

	"github.com/manleydev/golang-game-framework/rendering"
)

type Camera2D struct {
	Node2D
	Zoom     float64
	camera   *rendering.Camera2D
	renderer *rendering.Renderer2D
}

func (camera *Camera2D) Update(delta float64) error {
	if err := camera.Node2D.Update(delta); err != nil {
		return err
	}

	camRef := *(camera.camera)
	camRef.SetOffset(camera.GetGlobalPosition())
	camRef.SetZoom(camera.Zoom)

	return nil
}

func (camera *Camera2D) Enable() error {
	if camera.camera == nil {
		return errors.New("camera is nil")
	}
	return (*camera.renderer).SetCamera(camera.camera)
}

func (node *Camera2D) AddChild(child INode) {
	node.internalAddChild(node, child)
}

func NewCamera2D(renderer *rendering.Renderer2D, cam *rendering.Camera2D) *Camera2D {
	camera := Camera2D{
		Node2D:   NewNode2D(),
		camera:   cam,
		renderer: renderer,
		Zoom:     1.0,
	}

	return &camera
}
