package node

import (
	"github.com/manleydev/golang-game-framework/rendering"
)

type Sprite2D struct {
	Node2D
	texture rendering.Texture2D
	visible bool
}

func (sprite *Sprite2D) Draw(renderer rendering.Renderer2D) error {
	if !sprite.visible {
		return nil
	}
	if err := sprite.Node.Draw(renderer); err != nil {
		return err
	}

	return renderer.DrawTexture2D(rendering.DrawTexture2DInput{
		Texture:  sprite.texture,
		Rect:     sprite.texture.GetRect(),
		Position: sprite.GetGlobalPosition(),
		Rotation: sprite.Rotation,
	})
}

func (sprite *Sprite2D) SetVisible(v bool) {
	sprite.visible = v
}

func (node *Sprite2D) AddChild(child INode) {
	node.internalAddChild(node, child)
}

func NewSprite2D(texture rendering.Texture2D) *Sprite2D {
	sprite := Sprite2D{
		Node2D:  NewNode2D(),
		visible: true,
	}

	sprite.texture = texture

	return &sprite
}
